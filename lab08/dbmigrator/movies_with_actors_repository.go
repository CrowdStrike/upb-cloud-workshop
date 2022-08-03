package dbmigrator

import (
	"context"
	"database/sql"
	"lab08/domain"

	"github.com/alecthomas/log4go"
)

const (
	createTableStmt = `CREATE TABLE IF NOT EXISTS movies_with_actors(
							 id SERIAL PRIMARY KEY,
							 movie_id char(36),
							 actor_id char(36),
							 CONSTRAINT fk_movie
								FOREIGN KEY(movie_id)
									REFERENCES movies(id),
							 CONSTRAINT fk_actor
								 FOREIGN KEY(actor_id)
									 REFERENCES actors(id)
							);
						`
	selectInnerJoinStmt = `SELECT m.id, m.name, m.description, a.id, a.name, a.quote 
								from MOVIES AS m INNER JOIN movies_with_actors AS mwa on m.id=mwa.movie_id 
									INNER JOIN actors AS a ON mwa.actor_id=a.id 
							ORDER BY mwa.id
							LIMIT $1 OFFSET $2`
	insertBulkStmt = "INSERT INTO movies_with_actors (movie_id, actor_id) VALUES %s"
)

type movieWithActorRow struct {
	// TODO: task 3.b.: fill all the necesary fields to parse a row from the JOIN statement
}

type MoviesWithActorsRepository struct {
	db        *sql.DB
	movies    domain.Movies
	logger    *log4go.Logger
	batchSize int
}

func NewMoviesWithActorsRepository(
	db *sql.DB,
	movies domain.Movies,
	logger *log4go.Logger,
	batchSize int) *MoviesWithActorsRepository {

	mwa := MoviesWithActorsRepository{
		db:        db,
		movies:    movies,
		logger:    logger,
		batchSize: batchSize,
	}

	return &mwa
}

func (mwa *MoviesWithActorsRepository) MigrateMovies(ctx context.Context) error {
	// TODO: task 3.a. create the table movies_with_actors

	// TODO: task 3.a. do a full table scan over the movies table
	// hint: use the mwa.movies.ListMovies() call and fixed limit size
	// the stop condition should be, when we don't have any more records
	exit := true // TODO: remove this variable... it is just to illustrate that we do on application side, the full table scan
	for !exit {
		// TODO: task 3.a. retrieve a batch of records from the movies table

		// TODO: task 3.a. create batch insert statements for each pair ('movie_id', 'actor_id')

		// TODO: task 3.a. append all stored pairs ('movie_id', 'actor_id') (comma separated) and perform the operation
	}

	return nil
}

func (mwa *MoviesWithActorsRepository) ListAllMovies(ctx context.Context, limit, offset int) ([]domain.MovieExt, error) {
	// TODO: task 3.b.: run the select statement with the inner joins.

	// TODO: task 3.b.: use an intermediary struct to read all the rows

	// TODO: task 3.b.: emit the final array of objects

	return nil, nil
}
