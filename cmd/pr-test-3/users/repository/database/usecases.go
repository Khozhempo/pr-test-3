package database

import (
	"context"
)

func (r *RepositoryDbStore) CreateUser(ctx context.Context, user string) error {
	checkIfExist := UsersStruct{Username: user, Deleted: false}
	result := r.db.WithContext(ctx).Where(checkIfExist).First(&UsersStruct{})
	if result.RowsAffected != 0 {
		return ErrorUserAlreadyExist
	}

	userItemToCreate := UsersStruct{Username: user}
	result = r.db.WithContext(ctx).Create(&userItemToCreate)

	return result.Error
}

func (r *RepositoryDbStore) RemoveUser(ctx context.Context, user string) error {
	userItem := UsersStruct{
		Username: user,
		Deleted:  true}

	result := r.db.WithContext(ctx).Save(&userItem)

	return result.Error
}

func (r *RepositoryDbStore) GetAllUsers(ctx context.Context) ([]string, error) {
	var allUsers []string

	result := r.db.WithContext(ctx).Model(UsersStruct{}).Select("Username").Where("Deleted = ?", false).Find(&allUsers)
	if result.Error != nil {
		return nil, ErrorUnknownWhileQuery
	}

	return allUsers, nil
}
