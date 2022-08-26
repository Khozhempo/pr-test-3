package database

import (
	"gorm.io/gorm"
	"time"
)

type RepositoryDbStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *RepositoryDbStore {
	return &RepositoryDbStore{
		db: db,
	}
}

type UsersStruct struct {
	Username  string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool `gorm:"index;default:false"`
}
