package database

import "errors"

var (
	ErrorUserAlreadyExist  = errors.New("user already exists")
	ErrorUnknownWhileQuery = errors.New("error while querying list of all users")
)
