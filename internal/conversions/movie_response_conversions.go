package conversions

import (
	"time"

	"github.com/charlesoller/omni-import-microservice/internal/db"
	"github.com/charlesoller/omni-import-microservice/internal/models"
)

type movieResponseConverter struct {
	movie models.MovieDetailsResponse
}

func NewMovieResponseConverter(movie *models.MovieDetailsResponse) *movieResponseConverter {
	return &movieResponseConverter{}
}

func (s *movieResponseConverter)ToMovie() *db.UpsertMovieParams {
	m := s.movie

	releaseDate, _ := time.Parse("2006-01-02", m.ReleaseDate)

	var collectionID int32
	if m.Collection.ID != 0 {
		collectionID = int32(m.Collection.ID)
	}

	return &db.UpsertMovieParams{
		ID:               int32(m.ID),
		Title:            m.Title,
		OriginalTitle:    m.OriginalTitle,
		Overview:         m.Overview,
		ReleaseDate:      releaseDate,
		Runtime:          int32(m.Runtime),
		Budget:           m.Budget,
		Revenue:          m.Revenue,
		Popularity:       m.Popularity,
		VoteAverage:      m.VoteAverage,
		VoteCount:        int32(m.VoteCount),
		Status:           m.Status,
		Tagline:          m.Tagline,
		Homepage:         m.Homepage,
		OriginalLanguage: m.OriginalLanguage,
		Adult:            m.Adult,
		BackdropPath:     m.BackdropPath,
		PosterPath:       m.PosterPath,
		CollectionID:     collectionID,
	}
}

func (s *movieResponseConverter)ToCollection() *db.UpsertCollectionParams {
	m := s.movie

	return &db.UpsertCollectionParams{
		ID:           int32(m.Collection.ID),
		Name:         m.Title,     
		PosterPath:   m.PosterPath,
		BackdropPath: m.BackdropPath,
	}
}

func (s *movieResponseConverter)ToGenre() []*db.UpsertGenreParams {
	m := s.movie
	p := make([]*db.UpsertGenreParams, 0, len(m.Genres))

	for _, g := range m.Genres {
		p = append(p, &db.UpsertGenreParams{
			ID: g.ID,
			Name: g.Name,
		})
	}

	return p
}

func (s *movieResponseConverter)ToMovieGenre() []*db.UpsertMovieGenreParams {
	m := s.movie
	p := make([]*db.UpsertMovieGenreParams, 0, len(m.Genres))

	for _, g := range m.Genres {
		p = append(p, &db.UpsertMovieGenreParams{
			MovieID: int32(m.ID),
			GenreID: g.ID,
		})
	}

	return p
}