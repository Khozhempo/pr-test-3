package database

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	var err error
	ctx := context.Background()
	test := "username_test"

	err = godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := RepositoryDbStore{db: db}
	err = r.CreateUser(ctx, test)
	if err != nil {
		t.Error("response error: no expected", "received", err.Error())
	}

	list, err := r.GetAllUsers(ctx)
	if err != nil {
		t.Error("response error: no expected", "received", err.Error())
	}

	if !checkInStringSlice(list, test) {
		t.Error("no ", test, " in userbase")
	}
}

func TestRemoveUser(t *testing.T) {
	var err error
	ctx := context.Background()
	test := "username_test"

	err = godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := RepositoryDbStore{db: db}
	err = r.RemoveUser(ctx, test)
	if err != nil {
		t.Error("response error: no expected", "received", err.Error())
	}

	list, err := r.GetAllUsers(ctx)
	if err != nil {
		t.Error("response error: no expected", "received", err.Error())
	}

	if checkInStringSlice(list, test) {
		t.Error(test, " still in userbase")
	}
}

func TestGetAllUsers(t *testing.T) {
	var err error
	ctx := context.Background()

	err = godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := RepositoryDbStore{db: db}
	_, err = r.GetAllUsers(ctx)
	if err != nil {
		t.Error("response error: no expected", "received", err.Error())
	}
}

func checkInStringSlice(searchIn []string, lookFor string) bool {
	for _, value := range searchIn {
		if value == lookFor {
			return true
		}
	}
	return false
}
