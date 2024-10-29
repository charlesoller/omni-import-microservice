package services

import (
	"fmt"
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

func (s *tmdbService) makeSingleRequest(endpoint string) (*models.MovieDetailsResponse, error) {
	url := fmt.Sprintf("%s%s", s.baseUrl, endpoint)
	auth := fmt.Sprintf("Bearer %s", s.auth)

	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": auth,
	}

	movie, err := utils.MakeRequest[models.MovieDetailsResponse]("GET", url, headers, nil)
	return movie, err
}

func (s *tmdbService) makeMultipleRequest(endpoint string) ([]models.MovieDetailsResponse, error) {
	url := fmt.Sprintf("%s%s", s.baseUrl, endpoint)
	auth := fmt.Sprintf("Bearer %s", s.auth)

	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": auth,
	}

	type MultiResponse struct {
		Page         int                           `json:"page"`
		Results      []models.MovieDetailsResponse `json:"results"`
		TotalPages   int                           `json:"total_pages"`
		TotalResults int                           `json:"total_results"`
	}

	res, err := utils.MakeRequest[MultiResponse]("GET", url, headers, nil)

	return res.Results, err
}

func (s *tmdbService) GetMovieDetails(id int) (*models.MovieDetailsResponse, error) {
	endpoint := fmt.Sprintf("/movie/%s?append_to_response=credits", strconv.Itoa(id))
	movie, err := s.makeSingleRequest(endpoint)
	if err != nil || movie.ID == 0 {
		return nil, err
	}

	return movie, err
}

func (s *tmdbService) GetPopularMoviePageIds(id int) ([]int, error) {
	endpoint := fmt.Sprintf("/movie/popular?page=%s", strconv.Itoa(id))
	movies, err := s.makeMultipleRequest(endpoint)
	if err != nil {
		return nil, err
	}

	var ids []int
	for _, v := range movies {
		ids = append(ids, v.ID)
	}

	return ids, nil
}
