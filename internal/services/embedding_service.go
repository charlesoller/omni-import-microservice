package services

import (
	"github.com/charlesoller/omni-import-microservice/internal/models"
	"github.com/charlesoller/omni-import-microservice/internal/utils"
)

type embeddingService struct {
	url string
}

func NewEmbeddingService() *embeddingService {
	return &embeddingService{
		// Replace below with env
		url: "http://localhost:8000/api/embed",
	}
}

func (s *embeddingService) makeRequest(body *models.EmbeddingArg) (*models.EmbeddingResponse, error) {
	embedding, err := utils.MakeRequest[models.EmbeddingResponse]("POST", s.url, nil, *body)
	return embedding, err
}

func (s *embeddingService) EmbedMovie(m *models.EmbeddingArg) ([]float32, error) {
	embedding, err := s.makeRequest(m)
	if err != nil || embedding == nil {
		return nil, err
	}

	return embedding.Embedding, err
}
