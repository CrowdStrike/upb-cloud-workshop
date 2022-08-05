package queue

import "exam-api/domain"

// This part is bonus, if you finish everything else

type RedisRepo struct {
}

func NewRedisRepo() *RedisRepo {
	return &RedisRepo{}
}

func (r *RedisRepo) Add(batch []domain.Product) error {
	return nil
}
