package delivery

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"pr-test-3/users/repository/cache"
	"pr-test-3/users/repository/database"
	"pr-test-3/users/repository/kafka"
)

type Handler struct {
	Db    *database.RepositoryDbStore
	Cache *cache.RepositoryRedisStore
	Kafka *kafka.RepositoryKafkaStore
}

func NewHandler(db *database.RepositoryDbStore, redisConn *cache.RepositoryRedisStore, kafka *kafka.RepositoryKafkaStore) *Handler {
	return &Handler{
		Db:    db,
		Cache: redisConn,
		Kafka: kafka,
	}
}

func (u *Handler) AddUser(ctx context.Context, req *UserInfoRequest) (*UserResponse, error) {
	var err error

	err = u.Db.CreateUser(ctx, req.Username)
	if errors.Is(err, database.ErrorUserAlreadyExist) {
		err = status.Error(codes.AlreadyExists, err.Error())
		return &UserResponse{Success: false}, err
	}

	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		return &UserResponse{Success: false}, err
	}

	_, err = u.updateCache()
	if err != nil {
		log.Println("error while update cache", err.Error())
	}

	if err = u.Kafka.LogNewUser(req.Username); err != nil {
		log.Println("error while logging", err.Error())
	}

	return &UserResponse{Success: true}, status.Error(codes.OK, "Ok")
}

func (u *Handler) RemoveUser(ctx context.Context, req *UserInfoRequest) (*UserResponse, error) {
	var err error

	err = u.Db.RemoveUser(ctx, req.Username)
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		return &UserResponse{Success: false}, err
	}

	_, err = u.updateCache()
	if err != nil {
		log.Println("error while update cache", err.Error())
	}

	return &UserResponse{Success: true}, status.Error(codes.OK, "Ok")
}

func (u *Handler) ListAllUsers(ctx context.Context, req *NoArgs) (*ListOfUsersResponse, error) {
	var (
		err      error
		allUsers []string
	)

	allUsers, err = u.Cache.GetAllUsersFromCache()
	if err == nil || len(allUsers) != 0 {
		return &ListOfUsersResponse{Usernames: allUsers}, status.Error(codes.OK, "Ok")
	}

	if errors.Is(err, cache.ErrorCacheIsEmpty) || len(allUsers) == 0 {
		allUsers, err = u.updateCache()
		if err != nil {
			err = status.Error(codes.Unknown, err.Error())
			return &ListOfUsersResponse{}, err
		}

		return &ListOfUsersResponse{Usernames: allUsers}, status.Error(codes.OK, "Ok")
	}

	log.Printf("error while get data from cache: %s", err.Error())
	return &ListOfUsersResponse{}, status.Error(codes.NotFound, "error occurred, try again later")
}

func (u *Handler) mustEmbedUnimplementedUsersServiceServer() {
	return
}

func (u *Handler) updateCache() ([]string, error) {
	ctx := context.Background()
	allUsers, err := u.Db.GetAllUsers(ctx)
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		return nil, err
	}

	err = u.Cache.UpdateCache(allUsers)
	if err != nil {
		return nil, err
	}
	return allUsers, nil
}
