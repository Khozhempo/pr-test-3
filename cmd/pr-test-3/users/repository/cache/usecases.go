package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

type AllUsersInCache struct {
	Users []string `json:"allusers"`
}

var (
	redisKeyAllusers = "allusers"
	ctx              = context.Background()
)

func (r *RepositoryRedisStore) UpdateCache(allUsers []string) error {
	jsonToCache, err := json.Marshal(AllUsersInCache{Users: allUsers})
	if err != nil {
		return err
	}

	r.redisConn.Set(ctx, redisKeyAllusers, string(jsonToCache), r.timeToKeep)
	return nil
}

func (r *RepositoryRedisStore) GetAllUsersFromCache() ([]string, error) {
	var allUsers AllUsersInCache

	jsonFromCache, err := r.redisConn.Get(ctx, redisKeyAllusers).Result()
	if err == redis.Nil {
		return nil, ErrorCacheIsEmpty
	} else if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonFromCache), &allUsers)
	if err != nil {
		return nil, err
	}

	return allUsers.Users, nil
}
