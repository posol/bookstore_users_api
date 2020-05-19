package users

import (
	"fmt"
	"github.com/posol/bookstore_users_api/datasources/mysql/users_db"
	"github.com/posol/bookstore_users_api/utils/dates"
	"github.com/posol/bookstore_users_api/utils/errors"
	"strings"
)

const (
	queryInsertUser  = "insert into users(first_name, last_name, email, date_created) values(?, ?, ?, ?);"
	indexUniqueEmail = "email_uindex"
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

	user.DateCreated = dates.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewIntrenalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewIntrenalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId

	return nil
}
