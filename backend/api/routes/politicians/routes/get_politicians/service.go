package get_politicians

import "fmt"

type Service interface {
	getAndMapPoliticians() (ResponseDTO, error)
}

type service struct {
	repo Repository
}

func (s *service) getAndMapPoliticians() (ResponseDTO, error) {
	politicians, err := s.repo.getPoliticians()
	if err != nil {
		return ResponseDTO{}, fmt.Errorf("failed to get politician: %s", err)
	}

	respDTO := make(ResponseDTO, 0, len(politicians))

	for _, politician := range politicians {
		respDTO = append(respDTO, PoliticianDTO{
			ID:       politician.ID,
			FullName: politician.FirstName + " " + politician.LastName,
		})
	}

	return respDTO, nil
}
