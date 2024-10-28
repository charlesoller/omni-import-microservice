package conversions

import (
	"fmt"
	"time"
	"strings"
	"sort"

	"github.com/charlesoller/omni-import-microservice/internal/db"
	"github.com/charlesoller/omni-import-microservice/internal/models"
)

type MovieResponseConverter struct {
	movie *models.MovieDetailsResponse
}

func NewMovieResponseConverter(movie *models.MovieDetailsResponse) *MovieResponseConverter {
	return &MovieResponseConverter{
		movie: movie,
	}
}

func (s *MovieResponseConverter) ToMovie() *db.UpsertMovieParams {
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

func (s *MovieResponseConverter) ToCollection() *db.UpsertCollectionParams {
	m := s.movie

	return &db.UpsertCollectionParams{
		ID:           int32(m.Collection.ID),
		Name:         m.Title,
		PosterPath:   m.PosterPath,
		BackdropPath: m.BackdropPath,
	}
}

func (s *MovieResponseConverter) ToGenres() []*db.UpsertGenreParams {
	m := s.movie
	p := make([]*db.UpsertGenreParams, 0, len(m.Genres))

	for _, g := range m.Genres {
		p = append(p, &db.UpsertGenreParams{
			ID:   g.ID,
			Name: g.Name,
		})
	}

	return p
}

func (s *MovieResponseConverter) ToMovieGenres() []*db.UpsertMovieGenreParams {
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

func (s *MovieResponseConverter) ToProductionCompanies() []*db.UpsertProductionCompanyParams {
	m := s.movie
	p := make([]*db.UpsertProductionCompanyParams, 0, len(m.ProductionCompanies))

	for _, pc := range m.ProductionCompanies {
		p = append(p, &db.UpsertProductionCompanyParams{
			ID:            int32(pc.ID),
			Name:          pc.Name,
			LogoPath:      pc.LogoPath,
			OriginCountry: pc.OriginCountry,
		})
	}

	return p
}

func (s *MovieResponseConverter) ToMovieProductionCompanies() []*db.UpsertMovieProductionCompanyParams {
	m := s.movie
	p := make([]*db.UpsertMovieProductionCompanyParams, 0, len(m.ProductionCompanies))

	for _, pc := range m.ProductionCompanies {
		p = append(p, &db.UpsertMovieProductionCompanyParams{
			MovieID:   int32(m.ID),
			CompanyID: pc.ID,
		})
	}

	return p
}

func (s *MovieResponseConverter) ToCountries() []*db.UpsertCountryParams {
	m := s.movie
	p := make([]*db.UpsertCountryParams, 0, len(m.ProductionCountries))

	for _, c := range m.ProductionCountries {
		p = append(p, &db.UpsertCountryParams{
			Iso31661: c.Iso31661,
			Name:     c.Name,
		})
	}

	return p
}

func (s *MovieResponseConverter) ToMovieCountries() []*db.UpsertMovieCountryParams {
	m := s.movie
	p := make([]*db.UpsertMovieCountryParams, 0, len(m.ProductionCountries))

	for _, c := range m.ProductionCountries {
		p = append(p, &db.UpsertMovieCountryParams{
			MovieID:   int32(m.ID),
			CountryID: c.Iso31661,
		})
	}

	return p
}

func (s *MovieResponseConverter) ToLanguages() []*db.UpsertLanguageParams {
	m := s.movie
	p := make([]*db.UpsertLanguageParams, 0, len(m.Languages))

	for _, l := range m.Languages {
		p = append(p, &db.UpsertLanguageParams{
			Iso6391:     l.Iso6391,
			Name:        l.Name,
			EnglishName: l.EnglishName,
		})
	}

	return p
}

func (s *MovieResponseConverter) ToMovieLanguages() []*db.UpsertMovieLanguageParams {
	m := s.movie
	p := make([]*db.UpsertMovieLanguageParams, 0, len(m.Languages))

	for _, l := range m.Languages {
		p = append(p, &db.UpsertMovieLanguageParams{
			LanguageID: l.Iso6391,
			MovieID:    int32(m.ID),
		})
	}

	return p
}

func (s *MovieResponseConverter) ToEmbeddingArg() *models.EmbeddingArg {
	m := s.movie
	
	var genres []string
	for _, g := range m.Genres {
		genres = append(genres, g.Name)
	}

	var cast []models.CastMember
	cast = append(cast, m.Credits.Cast...)

	sort.Slice(cast, func(i, j int) bool {
		return cast[i].Order < cast[j].Order
	})

	limit := 5
	if len(cast) < limit {
		limit = len(cast)
	}

	var castNames []string
	for _, cam := range cast[:limit] {
		castNames = append(castNames, cam.Name)
	}

	var crew []models.CrewMember
	crew = append(crew, m.Credits.Crew...)

	sort.Slice(crew, func(i, j int) bool {
		return crew[i].Popularity < crew[j].Popularity
	})

	limit = 3
	if len(crew) < limit {
		limit = len(crew)
	}

	var crewNames []string
	for _, crm := range crew[:limit] {
		crewNames = append(crewNames, crm.Name)
	}

	genreString := strings.Join(genres, ", ")
	castString := strings.Join(castNames, ", ")
	crewString := strings.Join(crewNames, ", ")

	data := fmt.Sprintf("Title: %v\nOverview: %v\nRelease Date: %v\nGenres: %v\nCast: %v\nCrew: %v", m.Title, m.Overview, m.ReleaseDate, genreString, castString, crewString)

	return &models.EmbeddingArg{
		Data: data,
	}
}

func (s *MovieResponseConverter) ToCredits() int32 {
	m := s.movie
	return int32(m.ID)
}

func (s *MovieResponseConverter) ToCastMembers() []*db.UpsertCastMemberParams {
	c := s.movie.Credits.Cast
	p := make([]*db.UpsertCastMemberParams, 0, len(c))

	for _, cm := range c {
		p = append(p, &db.UpsertCastMemberParams{
			ID: int32(cm.ID),
			CastID: cm.CastID,
			CreditID: cm.CreditID,
			Gender: cm.Gender,
			Adult: cm.Adult,
			KnownForDepartment: cm.KnownForDepartment,
			Name: cm.Name,
			OriginalName: cm.OriginalName,
			Popularity: cm.Popularity,
			ProfilePath: cm.ProfilePath,
		})
	}

	return p
}

func (s *MovieResponseConverter) ToCrewMembers() []*db.UpsertCrewMemberParams {
	c := s.movie.Credits.Crew
	p := make([]*db.UpsertCrewMemberParams, 0, len(c))

	for _, cm := range c {
		p = append(p, &db.UpsertCrewMemberParams{
			ID: int32(cm.ID),
			CreditID: cm.CreditID,
			Gender: cm.Gender,
			Adult: cm.Adult,
			KnownForDepartment: cm.KnownForDepartment,
			Name: cm.Name,
			OriginalName: cm.OriginalName,
			Popularity: cm.Popularity,
			ProfilePath: cm.ProfilePath,
		})
	}

	return p
}


func (s *MovieResponseConverter) ToCreditsCastMembers() []*db.UpsertCreditCastMemberParams {
	m := s.movie
	c := m.Credits.Cast
	p := make([]*db.UpsertCreditCastMemberParams, 0, len(c))

	for _, cm := range c {
		p = append(p, &db.UpsertCreditCastMemberParams{
			CreditID: int32(m.ID),
			MemberID: cm.ID,
			Character: cm.Character,
			Order: int32(cm.Order),
		})
	}

	return p
}

func (s *MovieResponseConverter) ToCreditsCrewMembers() []*db.UpsertCreditCrewMemberParams {
	m := s.movie
	c := m.Credits.Crew
	p := make([]*db.UpsertCreditCrewMemberParams, 0, len(c))

	for _, cm := range c {
		p = append(p, &db.UpsertCreditCrewMemberParams{
			CreditID: int32(m.ID),
			MemberID: cm.ID,
			Job: cm.Job,
			Department: cm.Department,
		})
	}

	return p
}
