package cache

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"time"
)

func InitRedis() (*redis.Client, *time.Duration, error) {
	redisOptions, err := redisCreateDsn()
	if err != nil {
		return nil, nil, err
	}

	rdb := redis.NewClient(redisOptions)

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return nil, nil, errors.New("Error occurred while establishing connection to Redis " + err.Error())
	}

	keepInCache, err := redisGetKeepInCacheTime(os.Getenv("TEST_REDIS_TIMETOKEEP"))
	if err != nil {
		return nil, nil, err
	}

	return rdb, &keepInCache, nil
}

func redisCreateDsn() (*redis.Options, error) {
	host := os.Getenv("TEST_REDIS_HOST")
	if host == "" {
		return nil, errors.New("TEST_REDIS_HOST is not set")
	}

	port := os.Getenv("TEST_REDIS_PORT")
	if port == "" {
		return nil, errors.New("TEST_REDIS_PORT is not set")
	}

	db := os.Getenv("TEST_REDIS_DB")
	if db == "" {
		return nil, errors.New("TEST_REDIS_DB is not set")
	}

	dbInt, err := strconv.Atoi(db)
	if err != nil {
		return nil, errors.New("TEST_REDIS_DB set incorrect. Should be number")
	}

	return &redis.Options{
		Addr:     host + ":" + port,
		Password: os.Getenv("TEST_REDIS_PSW"),
		DB:       dbInt,
	}, nil
}

func redisGetKeepInCacheTime(str string) (time.Duration, error) {
	if str == "" {
		return 0, errors.New("necessary to set TEST_REDIS_TIMETOKEEP")
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New("incorrect number in TEST_REDIS_TIMETOKEEP")
	}

	return time.Second * time.Duration(i), nil
}
