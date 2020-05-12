package services

import (
	"github.com/posol/bookstore_users_api/domain/users"
	"github.com/posol/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
