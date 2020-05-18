package users

import (
	"fmt"
	"github.com/posol/bookstore_users_api/datasources/mysql/users_db"
	"github.com/posol/bookstore_users_api/utils/errors"
)

const (
	queryInsertUser = "insert into users(first_name, last_name, email, date_created) values(?, ?, ?, ?);"
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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewIntrenalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewIntrenalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewIntrenalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId

	/*current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = dates.GetNowString()

	usersDB[user.Id] = user*/
	return nil
}
