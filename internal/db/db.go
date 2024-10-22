// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.updateMovieEmbeddingStmt, err = db.PrepareContext(ctx, updateMovieEmbedding); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMovieEmbedding: %w", err)
	}
	if q.upsertCollectionStmt, err = db.PrepareContext(ctx, upsertCollection); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertCollection: %w", err)
	}
	if q.upsertCountryStmt, err = db.PrepareContext(ctx, upsertCountry); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertCountry: %w", err)
	}
	if q.upsertCountryISOStmt, err = db.PrepareContext(ctx, upsertCountryISO); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertCountryISO: %w", err)
	}
	if q.upsertGenreStmt, err = db.PrepareContext(ctx, upsertGenre); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertGenre: %w", err)
	}
	if q.upsertLanguageStmt, err = db.PrepareContext(ctx, upsertLanguage); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertLanguage: %w", err)
	}
	if q.upsertMovieStmt, err = db.PrepareContext(ctx, upsertMovie); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertMovie: %w", err)
	}
	if q.upsertMovieCountryStmt, err = db.PrepareContext(ctx, upsertMovieCountry); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertMovieCountry: %w", err)
	}
	if q.upsertMovieGenreStmt, err = db.PrepareContext(ctx, upsertMovieGenre); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertMovieGenre: %w", err)
	}
	if q.upsertMovieLanguageStmt, err = db.PrepareContext(ctx, upsertMovieLanguage); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertMovieLanguage: %w", err)
	}
	if q.upsertMovieProductionCompanyStmt, err = db.PrepareContext(ctx, upsertMovieProductionCompany); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertMovieProductionCompany: %w", err)
	}
	if q.upsertProductionCompanyStmt, err = db.PrepareContext(ctx, upsertProductionCompany); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertProductionCompany: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.updateMovieEmbeddingStmt != nil {
		if cerr := q.updateMovieEmbeddingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMovieEmbeddingStmt: %w", cerr)
		}
	}
	if q.upsertCollectionStmt != nil {
		if cerr := q.upsertCollectionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertCollectionStmt: %w", cerr)
		}
	}
	if q.upsertCountryStmt != nil {
		if cerr := q.upsertCountryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertCountryStmt: %w", cerr)
		}
	}
	if q.upsertCountryISOStmt != nil {
		if cerr := q.upsertCountryISOStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertCountryISOStmt: %w", cerr)
		}
	}
	if q.upsertGenreStmt != nil {
		if cerr := q.upsertGenreStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertGenreStmt: %w", cerr)
		}
	}
	if q.upsertLanguageStmt != nil {
		if cerr := q.upsertLanguageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertLanguageStmt: %w", cerr)
		}
	}
	if q.upsertMovieStmt != nil {
		if cerr := q.upsertMovieStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertMovieStmt: %w", cerr)
		}
	}
	if q.upsertMovieCountryStmt != nil {
		if cerr := q.upsertMovieCountryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertMovieCountryStmt: %w", cerr)
		}
	}
	if q.upsertMovieGenreStmt != nil {
		if cerr := q.upsertMovieGenreStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertMovieGenreStmt: %w", cerr)
		}
	}
	if q.upsertMovieLanguageStmt != nil {
		if cerr := q.upsertMovieLanguageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertMovieLanguageStmt: %w", cerr)
		}
	}
	if q.upsertMovieProductionCompanyStmt != nil {
		if cerr := q.upsertMovieProductionCompanyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertMovieProductionCompanyStmt: %w", cerr)
		}
	}
	if q.upsertProductionCompanyStmt != nil {
		if cerr := q.upsertProductionCompanyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertProductionCompanyStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                               DBTX
	tx                               *sql.Tx
	updateMovieEmbeddingStmt         *sql.Stmt
	upsertCollectionStmt             *sql.Stmt
	upsertCountryStmt                *sql.Stmt
	upsertCountryISOStmt             *sql.Stmt
	upsertGenreStmt                  *sql.Stmt
	upsertLanguageStmt               *sql.Stmt
	upsertMovieStmt                  *sql.Stmt
	upsertMovieCountryStmt           *sql.Stmt
	upsertMovieGenreStmt             *sql.Stmt
	upsertMovieLanguageStmt          *sql.Stmt
	upsertMovieProductionCompanyStmt *sql.Stmt
	upsertProductionCompanyStmt      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                               tx,
		tx:                               tx,
		updateMovieEmbeddingStmt:         q.updateMovieEmbeddingStmt,
		upsertCollectionStmt:             q.upsertCollectionStmt,
		upsertCountryStmt:                q.upsertCountryStmt,
		upsertCountryISOStmt:             q.upsertCountryISOStmt,
		upsertGenreStmt:                  q.upsertGenreStmt,
		upsertLanguageStmt:               q.upsertLanguageStmt,
		upsertMovieStmt:                  q.upsertMovieStmt,
		upsertMovieCountryStmt:           q.upsertMovieCountryStmt,
		upsertMovieGenreStmt:             q.upsertMovieGenreStmt,
		upsertMovieLanguageStmt:          q.upsertMovieLanguageStmt,
		upsertMovieProductionCompanyStmt: q.upsertMovieProductionCompanyStmt,
		upsertProductionCompanyStmt:      q.upsertProductionCompanyStmt,
	}
}
