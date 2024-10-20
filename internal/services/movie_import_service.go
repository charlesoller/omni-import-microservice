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
	ctx := context.Background()
	txCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	movie, err := s.tmdb.GetMovieDetails(id)
	if err != nil {
		// handle here
		log.Fatalln(err)
	}
	if movie == nil {
		// Don't try to convert etc if no movie found
		return
	}

	m := conversions.NewMovieResponseConverter(movie)

	err = s.db.ExecTx(txCtx, func(q *db.Queries) error {
		for _, v := range m.ToGenre() {
			_, err := q.UpsertGenre(txCtx, *v)
			if err != nil {
				log.Fatalln(err)
			}
		}
		_, err := q.UpsertCollection(txCtx, *m.ToCollection())
		if err != nil {
			log.Fatalln(err)
		}

		_, err = q.UpsertMovie(txCtx, *m.ToMovie())
		if err != nil {
			log.Fatalln(err)
		}

		for _, v := range m.ToMovieGenre() {
			_, err := q.UpsertMovieGenre(txCtx, *v)
			if err != nil {
				log.Fatalln(err)
			}
		}

		return err
	})

	if err != nil {
		// handle here
		log.Fatalln(err)
	}

	fmt.Printf("Successfully imported movie with id: %v\n", movie.Title)
}
