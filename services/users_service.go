package services

import (
	"github.com/posol/bookstore_users_api/domain/users"
	"github.com/posol/bookstore_users_api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	user := users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
