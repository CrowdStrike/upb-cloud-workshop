package sql

import (
	"context"
	"database/sql"
	"lab08/domain"

	"github.com/alecthomas/log4go"
	"github.com/lib/pq"
)

const (
	sqlCreateStmt = `INSERT INTO movies (id, name, description, actors)
					VALUES ($1, $2, $3, $4) 
					RETURNING id, name, description, actors`
	sqlDeleteStmt = `DELETE FROM movies WHERE id = $1 
					RETURNING id, name, description, actors`
	sqlGetByIDStmts = `SELECT id, name, description, actors
					FROM movies 
					WHERE id = ANY($1)`
	listMoviesStmt = `SELECT id, name, description, actors
				FROM movies
				ORDER BY id
				LIMIT $1 OFFSET $2`
)

type MovieRepository struct {
	db     *sql.DB
	logger *log4go.Logger
}

func NewMovieRepository(db *sql.DB, logger *log4go.Logger) *MovieRepository {
	mr := MovieRepository{
		db:     db,
		logger: logger,
	}

	return &mr
}

func (mr *MovieRepository) CreateMovie(ctx context.Context, movie domain.Movie) error {
	err := mr.db.QueryRowContext(ctx, sqlCreateStmt, movie.ID, movie.Name, movie.Description, pq.Array(movie.Actors))
	if err != nil && err.Err() != nil {
		return err.Err()
	}

	return nil
}

func (mr *MovieRepository) GetMoviesByID(ctx context.Context, ids []string) ([]domain.Movie, error) {
	rows, err := mr.db.QueryContext(ctx, sqlGetByIDStmts, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			_ = mr.logger.Error(err, domain.FailedRowsErrMsg)
		}
	}()

	var movies []domain.Movie
	for rows.Next() {
		movie := domain.Movie{}

		if err := rows.Scan(&movie.ID, &movie.Name, &movie.Description, pq.Array(&movie.Actors)); err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func (mr *MovieRepository) ListMovies(ctx context.Context, limit, offset int) ([]domain.Movie, error) {
	rows, err := mr.db.QueryContext(ctx, listMoviesStmt, limit, offset)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			_ = mr.logger.Error(err, domain.FailedRowsErrMsg)
		}
	}()

	var movies []domain.Movie
	for rows.Next() {
		var movie domain.Movie
		if err := rows.Scan(&movie.ID, &movie.Name, &movie.Description, pq.Array(&movie.Actors)); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (mr *MovieRepository) DeleteMovie(ctx context.Context, id string) error {
	rows, err := mr.db.QueryContext(ctx, sqlDeleteStmt, id)
	if err != nil {
		return err
	}
	if rows.Err() != nil {
		return rows.Err()
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			_ = mr.logger.Error(err, domain.FailedRowsErrMsg)
		}
	}()

	return nil
}
