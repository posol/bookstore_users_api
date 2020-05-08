package users

import (
	"fmt"

	"github.com/posol/bookstore_users_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user User) Get() *errors.RestError {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found ", user.Id))
	}

	user.Id = result.Id
	user.DateCreated = result.DateCreated
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName

	return nil
}

func (user User) Save() *errors.RestError {
	return nil
}
