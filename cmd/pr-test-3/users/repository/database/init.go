package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDB() (*gorm.DB, error) {
	dsn, err := dbCreateDsn()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Error occurred while establishing connection to PostgresDB: " + err.Error())
	}

	err = db.AutoMigrate(&UsersStruct{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func dbCreateDsn() (string, error) {
	host := os.Getenv("TEST_DB_HOST")
	if host == "" {
		return "", errors.New("TEST_DB_HOST is not set")
	}

	user := os.Getenv("TEST_DB_USER")
	if user == "" {
		return "", errors.New("TEST_DB_USER is not set")
	}

	password := os.Getenv("TEST_DB_PSW")
	if password == "" {
		return "", errors.New("TEST_DB_PSW is not set")
	}

	name := os.Getenv("TEST_DB_BASE")
	if name == "" {
		return "", errors.New("TEST_DB_BASE is not set")
	}

	port := os.Getenv("TEST_DB_PORT")
	if port == "" {
		return "", errors.New("TEST_DB_PORT is not set")
	}

	options := "sslmode=disable"

	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s",
		host, user, password, name, port, options), nil
}
