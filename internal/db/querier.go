// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	UpsertCollection(ctx context.Context, arg UpsertCollectionParams) (Collection, error)
	UpsertCountry(ctx context.Context, arg UpsertCountryParams) (Country, error)
	UpsertGenre(ctx context.Context, arg UpsertGenreParams) (Genre, error)
	UpsertMovie(ctx context.Context, arg UpsertMovieParams) (Movie, error)
	UpsertMovieGenre(ctx context.Context, arg UpsertMovieGenreParams) (MovieGenre, error)
	UpsertMovieProductionCompany(ctx context.Context, arg UpsertMovieProductionCompanyParams) (MovieProductionCompany, error)
	UpsertProductionCompany(ctx context.Context, arg UpsertProductionCompanyParams) (ProductionCompany, error)
}

var _ Querier = (*Queries)(nil)
