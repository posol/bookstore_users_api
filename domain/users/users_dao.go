package users

import (
	"fmt"

	"github.com/posol/bookstore_users_api/utils/dates"
	"github.com/posol/bookstore_users_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
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

func (user *User) Save() *errors.RestError {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = dates.GetNowString()

	usersDB[user.Id] = user
	return nil
}
