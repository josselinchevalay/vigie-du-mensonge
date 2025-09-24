package import_occupations

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"unicode"
	"vdm/data_import/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type nameKey struct{ first, last string }

// LoadFromCSV imports occupations.csv into occupations table (non-presidential entries).
// CSV format (comma-separated):
// id,gouvernement,code_fonction,prenom,nom,fonction,date_debut_fonction,date_fin_fonction
// - id is the Government reference_id (SMALLINT) used to resolve governments.id
// - Ensures a Politician exists for prenom/nom
// - Inserts an Occupation linked to the resolved Government, with presidential_reference NULL
// - Dates are parsed as YYYY-MM-DD; end_date may be empty (NULL)
func LoadFromCSV(db *gorm.DB) error {
	f, err := os.Open("occupations.csv")
	if err != nil {
		return fmt.Errorf("open occupations.csv: %w", err)
	}
	defer f.Close()

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ','
	r.FieldsPerRecord = -1 // allow optional last field

	// header
	if _, err := r.Read(); err != nil {
		return fmt.Errorf("read header: %w", err)
	}

	// caches: politician by name, governments by reference id
	polCache := make(map[nameKey]models.Politician)
	govCache := make(map[int]models.Government)

	const dateLayout = "2006-01-02"
	loc := time.UTC

	return db.Transaction(func(tx *gorm.DB) error {
		for {
			rec, err := r.Read()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				return fmt.Errorf("read record: %w", err)
			}
			if len(rec) < 8 {
				return fmt.Errorf("unexpected record length %d: %v", len(rec), rec)
			}

			// CSV columns:
			// 0: government reference id (SMALLINT)
			// 1: government name (ignored here)
			// 2: code_fonction
			// 3: prenom
			// 4: nom
			// 5: fonction (title)
			// 6: date_debut_fonction
			// 7: date_fin_fonction (may be empty)
			refStr := strings.TrimSpace(rec[0])
			code := strings.TrimSpace(rec[2])
			first := normalizeName(rec[3])
			last := normalizeName(rec[4])
			title := strings.TrimSpace(rec[5])
			startStr := strings.TrimSpace(rec[6])
			endStr := strings.TrimSpace(rec[7])

			ref, err := atoiStrict(refStr)
			if err != nil {
				return fmt.Errorf("parse government reference id '%s': %w", refStr, err)
			}

			start, err := time.ParseInLocation(dateLayout, startStr, loc)
			if err != nil {
				return fmt.Errorf("parse start date '%s': %w", startStr, err)
			}

			var end sql.NullTime
			if endStr != "" {
				te, err := time.ParseInLocation(dateLayout, endStr, loc)
				if err != nil {
					return fmt.Errorf("parse end date '%s': %w", endStr, err)
				}
				end = sql.NullTime{Valid: true, Time: te}
			}

			// ensure politician exists
			pol, err := getOrCreatePolitician(tx, polCache, first, last)
			if err != nil {
				return err
			}

			// resolve government by reference id
			gov, err := getGovernmentByRef(tx, govCache, ref)
			if err != nil {
				return err
			}

			// idempotency: check if an identical occupation already exists
			var existing models.Occupation
			query := tx.Where(
				"politician_id = ? AND government_id = ? AND presidential_reference IS NULL AND code = ? AND title = ? AND start_date = ? AND ((? AND end_date = ?) OR (? AND end_date IS NULL))",
				pol.ID, gov.ID, code, title, start,
				end.Valid, end.Time, !end.Valid,
			)
			switch err := query.First(&existing).Error; err {
			case nil:
				// already exists; skip
				continue
			case gorm.ErrRecordNotFound:
				// ok, proceed to create
			default:
				return fmt.Errorf("find existing occupation: %w", err)
			}

			occ := models.Occupation{
				PoliticianID:          pol.ID,
				GovernmentID:          uuidPtr(gov.ID),
				PresidentialReference: nil,
				Code:                  code,
				Title:                 title,
				StartDate:             start,
				EndDate:               end,
			}

			if err := tx.Create(&occ).Error; err != nil {
				return fmt.Errorf("insert occupation gov_ref=%d %s %s: %w", ref, first, last, err)
			}
		}
		return nil
	})
}

func getOrCreatePolitician(tx *gorm.DB, cache map[nameKey]models.Politician, first, last string) (models.Politician, error) {
	k := nameKey{first: first, last: last}
	if p, ok := cache[k]; ok {
		return p, nil
	}

	var p models.Politician
	if err := tx.Where("first_name = ? AND last_name = ?", first, last).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			p = models.Politician{FirstName: first, LastName: last}
			if err := tx.Create(&p).Error; err != nil {
				return models.Politician{}, fmt.Errorf("create politician %s %s: %w", first, last, err)
			}
		} else {
			return models.Politician{}, fmt.Errorf("find politician %s %s: %w", first, last, err)
		}
	}

	cache[k] = p
	return p, nil
}

func getGovernmentByRef(tx *gorm.DB, cache map[int]models.Government, ref int) (models.Government, error) {
	if g, ok := cache[ref]; ok {
		return g, nil
	}
	var g models.Government
	if err := tx.Where("reference = ?", ref).First(&g).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Government{}, fmt.Errorf("government with reference=%d not found; import governments first", ref)
		}
		return models.Government{}, fmt.Errorf("find government ref=%d: %w", ref, err)
	}
	cache[ref] = g
	return g, nil
}

func atoiStrict(s string) (int, error) {
	var n int
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, fmt.Errorf("non-digit in '%s'", s)
		}
		n = n*10 + int(r-'0')
	}
	return n, nil
}

// normalizeName trims and converts Unicode spaces (including NBSP) to normal spaces.
func normalizeName(s string) string {
	s = strings.TrimSpace(replaceAllSpaces(s))
	return s
}

func replaceAllSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || r == '\u00A0' { // NBSP
			return ' '
		}
		return r
	}, s)
}

func uuidPtr(id uuid.UUID) *uuid.UUID { return &id }
