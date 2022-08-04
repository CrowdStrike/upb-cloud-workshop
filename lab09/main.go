package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	"lab09/domain"
	"lab09/gateways/repositories"
)

const (
	localRedisHost   = ""
	defaultRedisPort = 6379
)

func main() {
	log := logrus.New()
	ctx := context.Background()

	redisURL := fmt.Sprintf("%s:%d", localRedisHost, defaultRedisPort)
	redisRepo := repositories.NewRedisRepoFromURL(redisURL, 0)

	rand.Seed(time.Now().Unix())
	timestamp := time.Now().Format(time.RFC3339Nano)
	randomSHA := fmt.Sprintf("%x", sha256.Sum256([]byte(timestamp)))

	metadata := &domain.FileMetadata{
		Sha256:          randomSHA,
		FileName:        "file123",
		CompilerVersion: int(rand.Int31n(200)),
		Size:            rand.Int63n(100000000000),
	}
	
	err := redisRepo.SaveFile(ctx, metadata)
	if err != nil {
		log.WithError(err).Fatal("Could not save data")
	}
	log.Infof("Successfully added metadata %+v", metadata)

	metadata, err = redisRepo.RetrieveFile(ctx, randomSHA)
	if err != nil {
		log.WithError(err).Fatal("Could not load data")
	}
	log.Infof("Successfully loaded metadata %+v", metadata)
	metadata, err = redisRepo.RetrieveFile(ctx, "this_does_not_exist")
	if err != nil {
		log.WithError(err).Fatal("Could not load data")
	}
	if metadata == nil {
		log.Infof("not found")
	}
}
