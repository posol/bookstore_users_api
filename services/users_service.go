package services

import (
	"github.com/posol/bookstore_users_api/domain/users"
	"github.com/posol/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}
