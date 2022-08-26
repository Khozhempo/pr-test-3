package cache

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type RepositoryRedisStore struct {
	redisConn  *redis.Client
	timeToKeep time.Duration
}

func NewStore(redisConn *redis.Client, duration time.Duration) *RepositoryRedisStore {
	return &RepositoryRedisStore{
		redisConn:  redisConn,
		timeToKeep: duration,
	}
}
