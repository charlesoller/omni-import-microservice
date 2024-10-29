package services

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/charlesoller/omni-import-microservice/internal/conversions"
	"github.com/charlesoller/omni-import-microservice/internal/database"
	"github.com/charlesoller/omni-import-microservice/internal/db"
	"github.com/pgvector/pgvector-go"
)

type movieImportService struct {
	tmdb  *tmdbService
	embed *embeddingService
	db    *database.Store
}

func NewMovieImportService(tmdb *tmdbService, embed *embeddingService, store *database.Store) *movieImportService {
	return &movieImportService{
		tmdb:  tmdb,
		embed: embed,
		db:    store,
	}
}

func (s *movieImportService) StartMultithreadedImport(numWorkers int, startIndex int, endIndex int) {
	numIndices := endIndex - startIndex + 1
	movieIndices := make(chan int, numIndices)

	fmt.Println("Loading movie indices...")
	for i := startIndex; i <= endIndex; i++ {
		movieIndices <- i
	}

	close(movieIndices)

	fmt.Printf("%v movie indices loaded!\n", numIndices)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	fmt.Printf("Wait group created with %v workers!\n", numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerId int) {
			defer wg.Done()
			for movieIndex := range movieIndices {
				log.Printf("Worker %d importing movie %d", workerId, movieIndex)
				s.importMovie(movieIndex)
				log.Printf("Worker %d finished importing movie %d", workerId, movieIndex)
			}
			log.Printf("Worker %d has no more movies to process", workerId)
		}(i)
	}

	wg.Wait() // Wait for all workers to finish
}

func (s *movieImportService) StartMultithreadedPopularImport(numWorkers int, startPage int, endPage int) {
	for page := startPage; page < endPage; page++ {
		fmt.Printf("Beginning import of page %v\n", page)

		ids, err := s.tmdb.GetPopularMoviePageIds(page)
		if err != nil || len(ids) == 0 {
			log.Fatalf("Failed to fetch movies on page: %v\nError: %v", page, err.Error())
		}

		fmt.Println("Loading movie indices...")
		movieIndices := make(chan int, len(ids))
		for _, v := range ids {
			movieIndices <- v
		}

		close(movieIndices)
		fmt.Printf("%v movie indices loaded!\n", len(ids))

		var wg sync.WaitGroup
		wg.Add(numWorkers)
		fmt.Printf("Wait group created with %v workers!\n", numWorkers)

		for i := 0; i < numWorkers; i++ {
			go func(workerId int) {
				defer wg.Done()
				for movieIndex := range movieIndices {
					log.Printf("Worker %d importing movie %d", workerId, movieIndex)
					s.importMovie(movieIndex)
					log.Printf("Worker %d finished importing movie %d", workerId, movieIndex)
				}
				log.Printf("Worker %d has no more movies to process", workerId)
			}(i)
		}
	
		wg.Wait() // Wait for all workers to finish
		log.Printf("Completed processing page %d", page)
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
		log.Fatalf("Failed to fetch movie with id: %d", id)
	}
	if movie == nil {
		// Don't try to convert etc if no movie found
		fmt.Printf("No movie found with ID: %v\n", id)
		return
}

	m := conversions.NewMovieResponseConverter(movie)

	if err = s.transact(m); err != nil {
		log.Fatalf("Failed to import movie with id: %d\nError: %v", id, err.Error())
	}
}

func (s *movieImportService) transact(m *conversions.MovieResponseConverter) error {
	ctx := context.Background()
	txCtx, cancel := context.WithTimeout(ctx, 300*time.Second)
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
		if err := s.upsertCredit(txCtx, m.ToCredits()); err != nil {
			return err
		}
		if err := s.upsertCastMembers(txCtx, m.ToCastMembers()); err != nil {
			return err
		}
		if err := s.upsertCrewMembers(txCtx, m.ToCrewMembers()); err != nil {
			return err
		}

		// // Join tables
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
		if err := s.upsertCreditsCastMembers(txCtx, m.ToCreditsCastMembers()); err != nil {
			return err
		}
		if err := s.upsertCreditsCrewMembers(txCtx, m.ToCreditsCrewMembers()); err != nil {
			return err
		}

		params, err := s.createEmbedding(m)
		if err != nil {
			return err
		}

		if err := s.updateEmbedding(txCtx, params); err != nil {
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

func (s *movieImportService) upsertCreditsCastMembers(txCtx context.Context, members []*db.UpsertCreditCastMemberParams) error {
	for _, m := range members {
		if _, err := s.db.Queries.UpsertCreditCastMember(txCtx, *m); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieImportService) upsertCreditsCrewMembers(txCtx context.Context, members []*db.UpsertCreditCrewMemberParams) error {
	for _, m := range members {
		if _, err := s.db.Queries.UpsertCreditCrewMember(txCtx, *m); err != nil {
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

func (s *movieImportService) upsertCredit(txCtx context.Context, id int32) error {
	if err := s.db.Queries.UpsertCredit(txCtx, id); err != nil {
		return err
	}

	return nil
}

func (s *movieImportService) upsertCastMembers(txCtx context.Context, members []*db.UpsertCastMemberParams) error {
	for _, m := range members {
		if _, err := s.db.Queries.UpsertCastMember(txCtx, *m); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieImportService) upsertCrewMembers(txCtx context.Context, members []*db.UpsertCrewMemberParams) error {
	for _, m := range members {
		if _, err := s.db.Queries.UpsertCrewMember(txCtx, *m); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieImportService) updateEmbedding(txCtx context.Context, arg *db.UpdateMovieEmbeddingParams) error {
	if err := s.db.Queries.UpdateMovieEmbedding(txCtx, *arg); err != nil {
		return err
	}

	return nil
}

func (s *movieImportService) createEmbedding(m *conversions.MovieResponseConverter) (*db.UpdateMovieEmbeddingParams, error) {
	embedding, err := s.embed.EmbedMovie(m.ToEmbeddingArg())

	// fmt.Printf("Type: %T \t Length: %v\n", *embedding, len(*embedding))
	if err != nil || embedding == nil {
		return nil, err
	}

	upsertParams := db.UpdateMovieEmbeddingParams{
		ID:        m.ToMovie().ID,
		Embedding: pgvector.NewVector(*embedding),
	}

	return &upsertParams, nil
}
