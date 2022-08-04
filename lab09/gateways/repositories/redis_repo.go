package repositories

import (
	"lab09/domain"
	"context"
	"time"
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

const queueKey = "lab09:file_queue"
const queueSizeKey = "lab09:file_queue:size"
const (
	addToQueueScript = "if redis.call('rpush', KEYS[1], ARGV[1]) > 0 then " +
		"return redis.call('incrby', KEYS[2], ARGV[2]) end"
)

var log = logrus.New()

type RedisRepo struct {
	redisPool *redis.Pool
	dbNr        int
}

var _ domain.FilesMetadataStorage = (*RedisRepo)(nil)


func NewRedisRepo(redisPool *redis.Pool, dbNr int) *RedisRepo {
	return &RedisRepo{
		redisPool: redisPool,
		dbNr:      dbNr,
	}
}

func NewRedisRepoFromURL(redisURL string, dbNr int) *RedisRepo {
	return &RedisRepo{
		redisPool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
	
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", redisURL) // we could also have timeouts here
				if err != nil {
					return nil, err
				}
				return c, err
			},
		},
		dbNr:      dbNr,
	}
}


func (r *RedisRepo) SaveFile(ctx context.Context, metadata *domain.FileMetadata) error {
    conn := r.redisPool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Warn("Redis connection didn't close properly")
		}
	}()
	// Select the db space for the command
	err := conn.Send("SELECT", r.dbNr)
	if err != nil {
		return fmt.Errorf("failed to select database %w", err)
	}

	marshalledMetadata, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("failed to marshall metadata %w", err)
	}
	err = conn.Send("SET", metadata.Sha256, marshalledMetadata)
	if err != nil {
		return fmt.Errorf("failed to save metadata in redis %w", err)
	}

	return nil
}

func (r *RedisRepo) RetrieveFile(ctx context.Context, sha256 string) (*domain.FileMetadata, error) {
    conn := r.redisPool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.WithError(err).Warn("Redis connection didn't close properly")
		}
	}()
	// Select the db space for the command
	err := conn.Send("SELECT", r.dbNr)
	if err != nil {
		return nil, fmt.Errorf("failed to select database %w", err)
	}

	var metadata domain.FileMetadata
	marshalledMeta, err := redis.Bytes(conn.Do("GET", sha256))
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get metadata from redis %w", err)
	}
	err = json.Unmarshal(marshalledMeta, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshall metadata %w", err)
	}

	return &metadata, nil
}

func (r *RedisRepo) AddFileToQueue(ctx context.Context, metadata *domain.FileMetadata) error {
	conn := r.redisPool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.WithError(err).Warn("Redis connection didn't close properly")
		}
	}()
	// Select the db space for the command
	err := conn.Send("SELECT", r.dbNr)
	if err != nil {
		return fmt.Errorf("failed to select database %w", err)
	}

	err = conn.Send("EVAL", addToQueueScript, 2, queueKey, queueSizeKey, metadata.Sha256, metadata.Size)
	if err != nil {
		return fmt.Errorf("failed to save metadata in redis %w", err)
	}

	return nil
}
