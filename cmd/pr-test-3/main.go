package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"pr-test-3/server"
)

func main() {
	// инициализация сервиса
	app, err := server.NewApp()
	if err != nil {
		log.Fatalf("Can't start, have errors: %s", err.Error())
	}

	// запуск сервера
	if err := app.Run(os.Getenv("TEST_PORT")); err != nil {
		log.Fatalf("%s", err.Error())
	}

}

func init() {
	// загрузка конфигурации: из памяти или .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
}
