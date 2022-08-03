package sql

import (
	"context"
	"database/sql"
	"lab08/domain"

	"github.com/alecthomas/log4go"
)

const (
	sqlActorsCreateStmt = `INSERT INTO actors (id, name, quote)
					VALUES ($1, $2, $3) 
					RETURNING id, name, quote`
	sqlActorsDeleteStmt = `DELETE FROM actors WHERE id = $1 
					RETURNING id, name, quote`
	sqlActorsGetByIDStmt = `SELECT id, name, quote
					FROM actors 
					WHERE id = ANY ($1)`
	listActorsStmt = `SELECT id, name, quote
				FROM actors 
				ORDER BY id
				LIMIT $1 OFFSET $2`
	insertActorsBulkStmt        = "INSERT INTO actors (id, name, quote) VALUES %s"
	countActorsPartialMatchStmt = "SELECT count(*) from actors where name like"
	countActorsExactMatchStmt   = "SELECT count(*) from actors where name="
)

type ActorRepository struct {
	db     *sql.DB
	logger *log4go.Logger
}

func NewActorRepository(db *sql.DB, logger *log4go.Logger) *ActorRepository {
	mr := ActorRepository{
		db:     db,
		logger: logger,
	}

	return &mr
}

func (mr *ActorRepository) CreateActor(ctx context.Context, actor domain.Actor) error {
	// TODO: implement me
	return nil
}

func (mr *ActorRepository) CreateActors(ctx context.Context, actors []domain.Actor) error {
	// TODO: task 5: index bulk data in the actors table
	return nil
}

func (mr *ActorRepository) GetActorsByID(ctx context.Context, ids []string) ([]domain.Actor, error) {
	// TODO: implement me
	return nil, nil
}

func (mr *ActorRepository) ListActors(ctx context.Context, limit, offset int) ([]domain.Actor, error) {
	// TODO: implement me
	return nil, nil
}

func (mr *ActorRepository) DeleteActor(ctx context.Context, id string) error {
	// TODO: implement me
	return nil
}

func (mr *ActorRepository) Count(ctx context.Context, name string, useExactMatch bool) (int, error) {
	var stmt string
	if useExactMatch {
		stmt = countActorsExactMatchStmt + "'" + name + "'"
	} else {
		stmt = countActorsPartialMatchStmt + "'%" + name + "%'"
	}

	// TODO: also remove this log line, as it will get quite spammy
	mr.logger.Info("Running count statement: %s", stmt)
	// TODO: task 5.b.:

	return 0, nil
}