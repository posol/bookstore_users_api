package users

import (
	"github.com/posol/bookstore_users_api/datasources/mysql/users_db"
	"github.com/posol/bookstore_users_api/utils/dates"
	"github.com/posol/bookstore_users_api/utils/errors"
	"github.com/posol/bookstore_users_api/utils/mysql_errors"
)

const (
	indexUniqueEmail = "email_uindex"
	queryInsertUser  = "insert into users(first_name, last_name, email, date_created) values(?, ?, ?, ?);"
	queryGetUserById = "select id, first_name, last_name, email, date_created from users where id = ?;"
	queryUpdateUser  = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
)

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewIntrenalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_errors.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewIntrenalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = dates.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_errors.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_errors.ParseError(err)
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewIntrenalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_errors.ParseError(err)
	}
	return nil
}
