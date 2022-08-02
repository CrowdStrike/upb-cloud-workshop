package domain

import (
	"context"
	"fmt"
	"time"
)

const FailedRowsErrMsg = "failed to close the rows"

// Movie ...
type Movie struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Actors      []string `json:"actors"`
}

// String ...
func (m Movie) String() string {
	return fmt.Sprintf("%#v", m)
}

// Movies ...
type Movies interface {
	CreateMovie(ctx context.Context, movie Movie) error

	GetMoviesByID(ctx context.Context, ids []string) ([]Movie, error)
	ListMovies(ctx context.Context, limit, offset int) ([]Movie, error)

	DeleteMovie(ctx context.Context, id string) error
}

// Actor ...
type Actor struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Quote string `json:"quote"`
}

// Actors ...
type Actors interface {
	CreateActor(ctx context.Context, actor Actor) error
	CreateActors(ctx context.Context, actors []Actor) error

	GetActorsByID(ctx context.Context, ids []string) ([]Actor, error)
	ListActors(ctx context.Context, limit, offset int) ([]Actor, error)

	DeleteActor(ctx context.Context, id string) error

	Count(ctx context.Context, name string, useExactMatch bool) (int, error)
}

// MovieExt ...
type MovieExt struct {
	Name        string
	Description string
	Actors      []Actor
}

type ExtendedMovies interface {
	MigrateMovies(ctx context.Context) error
	ListAllMovies(ctx context.Context, limit, offset int) ([]MovieExt, error)
}

type GeneratorCounter struct {
	PartialMatchName     string
	PartialMatchCount    int
	PartialMatchDuration time.Duration
	ExactMatchName       string
	ExactMatchCount      int
	ExactMatchDuration   time.Duration
}

type Generator interface {
	Generate(ctx context.Context, numRecords int) error
	Count(ctx context.Context) (GeneratorCounter, error)
}
