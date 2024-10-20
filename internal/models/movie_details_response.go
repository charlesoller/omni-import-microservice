package models

import (
	"github.com/charlesoller/omni-import-microservice/internal/db"
)

type languagePartial struct {
	EnglishName string `json:"english_name"`
	Iso6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type countryPartial struct {
	Iso31661 string `json:"iso_3166_1"`
	Name     string `json:"name"`
}

type MovieDetailsResponse struct {
	ID                  int                    `json:"id" db:"id"`
	Adult               bool                   `json:"adult"`
	BackdropPath        string                 `json:"backdrop_path"`
	Collection          db.Collection          `json:"belongs_to_collection"`
	Budget              int64                  `json:"budget"`
	Genres              []db.Genre             `json:"genres"`
	Homepage            string                 `json:"homepage"`
	ImdbId              string                 `json:"imdb_id"`
	OriginCountry       []string               `json:"origin_country"`
	OriginalLanguage    string                 `json:"original_language"`
	OriginalTitle       string                 `json:"original_title"`
	Overview            string                 `json:"overview"`
	Popularity          float64                `json:"popularity"`
	PosterPath          string                 `json:"poster_path"`
	ProductionCompanies []db.ProductionCompany `json:"production_companies"`
	ProductionCountries []countryPartial       `json:"production_countries"`
	ReleaseDate         string                 `json:"release_date"`
	Revenue             int64                  `json:"revenue"`
	Runtime             int                    `json:"runtime"`
	Languages           []languagePartial      `json:"spoken_language"`
	Status              string                 `json:"status"`
	Tagline             string                 `json:"tagline"`
	Title               string                 `json:"title"`
	VoteAverage         float64                `json:"vote_average"`
	VoteCount           int                    `json:"vote_count"`
}
