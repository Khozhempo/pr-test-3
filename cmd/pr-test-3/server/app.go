package server

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"os/signal"
	"pr-test-3/users/delivery"
	"pr-test-3/users/repository/cache"
	"pr-test-3/users/repository/database"
	"pr-test-3/users/repository/kafka"
	"time"
)

type AppStruct struct {
	db          *gorm.DB
	redisConn   *redis.Client
	keepInCache *time.Duration
	kafkaConn   *sarama.AsyncProducer
	kafkaTopic  *string
} // AppStruct конфигурация сервиса

func NewApp() (*AppStruct, error) {
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	redisConn, keepInCache, err := cache.InitRedis()
	if err != nil {
		return nil, err
	}

	kafkaConn, kafkaTopic, err := kafka.InitKafka()
	if err != nil {
		return nil, err
	}

	return &AppStruct{
		db:          db,
		redisConn:   redisConn,
		keepInCache: keepInCache,
		kafkaConn:   kafkaConn,
		kafkaTopic:  kafkaTopic,
	}, nil
} // NewApp инициализация конфигурации

func (a *AppStruct) Run(port string) error {
	server := grpc.NewServer()
	dbConn := database.NewStore(a.db)

	redisConn := cache.NewStore(a.redisConn, *a.keepInCache)
	defer a.redisConn.Close()

	kafkaConn := kafka.NewStore(a.kafkaConn, a.kafkaTopic)
	kafkaConn.StartListener()

	handler := delivery.NewHandler(dbConn, redisConn, kafkaConn)
	delivery.RegisterUsersServiceServer(server, handler)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return errors.New("Unable to create grpc listener: " + err.Error())
	}

	go func() {
		if err = server.Serve(listener); err != nil {
			log.Fatalf("Unable to start server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	log.Println("Start working...")
	<-quit

	log.Println("Got shutdown signal. Waiting before stop.")
	time.Sleep(5 * time.Second)

	server.Stop()
	return nil
} // Run запуск сервера
