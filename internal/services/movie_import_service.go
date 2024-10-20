package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/charlesoller/omni-import-microservice/internal/conversions"
	"github.com/charlesoller/omni-import-microservice/internal/database"
	"github.com/charlesoller/omni-import-microservice/internal/db"
)

type movieImportService struct {
	tmdb *tmdbService
	db   *database.Store
}

func NewMovieImportService(tmdb *tmdbService, store *database.Store) *movieImportService {
	return &movieImportService{
		tmdb: tmdb,
		db:   store,
	}
}

func (s *movieImportService) StartImport(i int) {
	for {
		s.importMovie(i)
		i++
	}
}

func (s *movieImportService) importMovie(id int) {
	movie, err := s.tmdb.GetMovieDetails(id)
	if err != nil {
		// handle here
		log.Fatalln(err)
	}
	if movie == nil {
		// Don't try to convert etc if no movie found
		fmt.Printf("No movie found with ID: %v\n", id)
		return
	}

	m := conversions.NewMovieResponseConverter(movie)

	if err = s.transact(m); err != nil {
		// handle here
		log.Fatalln(err)
	}

	fmt.Printf("Successfully imported movie!\t%v\t%v\n", movie.ID, movie.Title)
}

func (s *movieImportService) transact(m *conversions.MovieResponseConverter) error {
	ctx := context.Background()
	txCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := s.db.ExecTx(txCtx, func(q *db.Queries) error {
		if err := s.upsertCountries(txCtx, m.ToCountries()); err != nil {
			return err
		}
		if err := s.upsertLanguages(txCtx, m.ToLanguages()); err != nil {
			return err
		}
		if err := s.upsertGenres(txCtx, m.ToGenres()); err != nil {
			return err
		}
		if err := s.upsertProductionCompanies(txCtx, m.ToProductionCompanies()); err != nil {
			return err
		}
		if err := s.upsertCollection(txCtx, m.ToCollection()); err != nil {
			return err
		}
		if err := s.upsertMovie(txCtx, m.ToMovie()); err != nil {
			return err
		}
		if err := s.upsertMovieGenres(txCtx, m.ToMovieGenres()); err != nil {
			return err
		}
		if err := s.upsertMovieProductionCompanies(txCtx, m.ToMovieProductionCompanies()); err != nil {
			return err
		}
		if err := s.upsertMovieCountries(txCtx, m.ToMovieCountries()); err != nil {
			return err
		}
		if err := s.upsertMovieLanguages(txCtx, m.ToMovieLanguages()); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *movieImportService) upsertCountries(txCtx context.Context, countries []*db.UpsertCountryParams) error {
	for _, c := range countries {
		if _, err := s.db.Queries.UpsertCountry(txCtx, *c); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieImportService) upsertLanguages(txCtx context.Context, languages []*db.UpsertLanguageParams) error {
	for _, l := range languages {
		if _, err := s.db.Queries.UpsertLanguage(txCtx, *l); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieImportService) upsertGenres(txCtx context.Context, genres []*db.UpsertGenreParams) error {
	for _, g := range genres {
		if _, err := s.db.Queries.UpsertGenre(txCtx, *g); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieImportService) upsertProductionCompanies(txCtx context.Context, pcs []*db.UpsertProductionCompanyParams) error {
	for _, pc := range pcs {
		s.db.Queries.UpsertCountryISO(txCtx, pc.OriginCountry)
		if _, err := s.db.Queries.UpsertProductionCompany(txCtx, *pc); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieImportService) upsertMovieGenres(txCtx context.Context, genres []*db.UpsertMovieGenreParams) error {
	for _, g := range genres {
		if _, err := s.db.Queries.UpsertMovieGenre(txCtx, *g); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieImportService) upsertMovieProductionCompanies(txCtx context.Context, pcs []*db.UpsertMovieProductionCompanyParams) error {
	for _, pc := range pcs {
		if _, err := s.db.Queries.UpsertMovieProductionCompany(txCtx, *pc); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieImportService) upsertMovieCountries(txCtx context.Context, countries []*db.UpsertMovieCountryParams) error {
	for _, c := range countries {
		if _, err := s.db.Queries.UpsertMovieCountry(txCtx, *c); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieImportService) upsertMovieLanguages(txCtx context.Context, languages []*db.UpsertMovieLanguageParams) error {
	for _, l := range languages {
		if _, err := s.db.Queries.UpsertMovieLanguage(txCtx, *l); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieImportService) upsertCollection(txCtx context.Context, c *db.UpsertCollectionParams) error {
	_, err := s.db.Queries.UpsertCollection(txCtx, *c)
	if err != nil {
		return err
	}
	return nil
}

func (s *movieImportService) upsertMovie(txCtx context.Context, m *db.UpsertMovieParams) error {
	_, err := s.db.Queries.UpsertMovie(txCtx, *m)
	if err != nil {
		return err
	}
	return nil
}
