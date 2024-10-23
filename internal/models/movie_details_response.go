package models

import (
	"github.com/charlesoller/omni-import-microservice/internal/db"
)

type CastMember struct {
	db.CastMember
	Character string `json:"character"`
	Order     int    `json:"order"`
}

type CrewMember struct {
	db.CrewMember
	Job        string `json:"job"`
	Department string `json:"department"`
}

type CreditsResponse struct {
	Cast []CastMember `json:"cast"`
	Crew []CrewMember `json:"crew"`
}
type MovieDetailsResponse struct {
	ID                  int                    `json:"id"`
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
	ProductionCountries []db.Country           `json:"production_countries"`
	ReleaseDate         string                 `json:"release_date"`
	Revenue             int64                  `json:"revenue"`
	Runtime             int                    `json:"runtime"`
	Languages           []db.Language          `json:"spoken_languages"`
	Status              string                 `json:"status"`
	Tagline             string                 `json:"tagline"`
	Title               string                 `json:"title"`
	VoteAverage         float64                `json:"vote_average"`
	VoteCount           int                    `json:"vote_count"`
	Credits             CreditsResponse        `json:"credits"`
}
