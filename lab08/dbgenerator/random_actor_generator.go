package dbgenerator

import (
	"context"
	"fmt"
	"lab08/domain"
	"math/rand"

	"github.com/alecthomas/log4go"
	"github.com/google/uuid"
)

var firstNames = []string{
	"Liam",
	"Noah",
	"Oliver",
	"Elijah",
	"Olivia",
	"Emma",
	"Charlotte",
	"Amelia",
}

var lastNames = []string{
	"Smith",
	"Johnson",
	"Williams",
	"Brown",
	"Jones",
	"Miller",
	"Davis",
	"Garcia",
}

type generatedResult struct {
	succesfulCount int
	failedCount    int
	err            error
}

type RandomActorGenerator struct {
	actors     domain.Actors
	logger     *log4go.Logger
	numWorkers int
	batchSize  int
}

func NewRandomActorGenerator(
	actors domain.Actors,
	logger *log4go.Logger,
	numWorkers, batchSize int) *RandomActorGenerator {

	rg := RandomActorGenerator{
		actors:     actors,
		logger:     logger,
		numWorkers: numWorkers,
		batchSize:  batchSize,
	}

	return &rg
}

func (rg *RandomActorGenerator) Generate(ctx context.Context, numRecords int) error {
	// TODO: task 4.a.: generate bulk random data, accounting number of workers, batch size and number of records

	return nil
}

func (rg *RandomActorGenerator) Count(ctx context.Context) (domain.GeneratorCounter, error) {
	// TODO: task 5.a.: profile the count operations for both exact and partial matches

	return domain.GeneratorCounter{}, nil
}

func (rg *RandomActorGenerator) createNewActor() domain.Actor {
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]

	a := domain.Actor{
		Name:  fmt.Sprintf("%s %s", firstName, lastName),
		Quote: "Lorem Ipsum",
		ID:    uuid.New().String(),
	}

	return a
}
