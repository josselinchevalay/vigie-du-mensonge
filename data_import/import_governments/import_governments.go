package import_governments

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

	"gorm.io/gorm"
)

type key struct{ first, last string }

// LoadFromCSV imports governments.csv into politician and governments tables.
// CSV format (comma-separated):
// id,prenom,nom,date_debut_fonction,date_fin_fonction,duree
// - Ensures a Politician exists for prenom/nom
// - Inserts a Government with reference set to id
// - Dates are parsed as YYYY-MM-DD; end_date may be empty (NULL)
func LoadFromCSV(db *gorm.DB) error {
	f, err := os.Open("governments.csv")
	if err != nil {
		return fmt.Errorf("open governments.csv: %w", err)
	}
	defer f.Close()

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ','
	r.FieldsPerRecord = -1 // allow empty last column

	// header
	if _, err := r.Read(); err != nil {
		return fmt.Errorf("read header: %w", err)
	}

	// cache politician by normalized first/last
	cache := make(map[key]models.Politician)

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
			if len(rec) < 6 {
				return fmt.Errorf("unexpected record length %d: %v", len(rec), rec)
			}

			refStr := strings.TrimSpace(rec[0])
			first := normalizeName(rec[1])
			last := normalizeName(rec[2])
			startStr := strings.TrimSpace(rec[3])
			endStr := strings.TrimSpace(rec[4]) // may be empty

			ref, err := atoiStrict(refStr)
			if err != nil {
				return fmt.Errorf("parse reference id '%s': %w", refStr, err)
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

			pol, err := getOrCreatePolitician(tx, cache, first, last)
			if err != nil {
				return err
			}

			gov := models.Government{
				PrimeMinisterID: pol.ID,
				Reference:       ref,
				StartDate:       start,
				EndDate:         end,
			}

			if err := tx.Create(&gov).Error; err != nil {
				if isUniqueViolation(err) {
					// already inserted — skip
					continue
				}
				return fmt.Errorf("insert government ref=%d: %w", ref, err)
			}
		}
		return nil
	})
}

func getOrCreatePolitician(tx *gorm.DB, cache map[key]models.Politician, first, last string) (models.Politician, error) {
	k := key{first: first, last: last}
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

func isUniqueViolation(err error) bool {
	e := strings.ToLower(err.Error())
	return strings.Contains(e, "duplicate key") || strings.Contains(e, "23505")
}
