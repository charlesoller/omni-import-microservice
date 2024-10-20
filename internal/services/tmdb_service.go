package services

import (
	"fmt"
	"log"
	"strconv"

	"github.com/charlesoller/omni-import-microservice/internal/models"
	"github.com/charlesoller/omni-import-microservice/internal/utils"
)

type tmdbService struct {
	baseUrl string
	auth    string
}

func NewTmdbService(auth string) *tmdbService {
	return &tmdbService{
		baseUrl: "https://api.themoviedb.org/3",
		auth:    auth,
	}
}

func (s *tmdbService) makeRequest(endpoint string) (*models.MovieDetailsResponse, error) {
	url := fmt.Sprintf("%s%s", s.baseUrl, endpoint)
	auth := fmt.Sprintf("Bearer %s", s.auth)

	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": auth,
	}

	movie, err := utils.MakeRequest[models.MovieDetailsResponse](url, headers)
	return movie, err
}

func (s *tmdbService) GetMovieDetails(id int) (*models.MovieDetailsResponse, error) {
	endpoint := fmt.Sprintf("/movie/%s", strconv.Itoa(id))
	movie, err := s.makeRequest(endpoint)
	if err != nil {
		// handle error better here
		log.Fatalln(err)
	}
	if movie.ID == 0 {
		fmt.Printf("No movie found with id: %v\n", id)
		return nil, err
	}

	return movie, err
}