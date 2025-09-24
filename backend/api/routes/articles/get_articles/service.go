package get_articles

type Service interface {
	getAndMapArticles() (ResponseDTO, error)
}

type service struct {
	repo Repository
}

func (s *service) getAndMapArticles() (ResponseDTO, error) {
	articles, err := s.repo.getArticles()
	if err != nil {
		return ResponseDTO{}, err
	}

	respDTO := make(ResponseDTO, len(articles))

	for i := range articles {
		respDTO[i] = newArticleDTO(articles[i])
	}

	return respDTO, nil
}
