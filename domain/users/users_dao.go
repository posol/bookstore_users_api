package users

import (
	"fmt"

	"github.com/posol/bookstore_users_api/datasources/mysql/users_db"
	"github.com/posol/bookstore_users_api/logger"
	"github.com/posol/bookstore_users_api/utils/errors"
)

const (
	queryInsertUser       = "insert into users(first_name, last_name, email, date_created, password, status) values(?, ?, ?, ?, ?, ?);"
	queryGetUserById      = "select id, first_name, last_name, email, date_created, status from users where id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser       = "delete from users where id=?;"
	queryFindUserByStatus = "select id, first_name, last_name, email, date_created, status from users where status=?;"
)

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewIntrenalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return errors.NewIntrenalServerError("database error")
		//return mysql_errors.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewIntrenalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.NewIntrenalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return errors.NewIntrenalServerError("database error")
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewIntrenalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewIntrenalServerError("database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewIntrenalServerError("database error")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewIntrenalServerError("database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, errors.NewIntrenalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, errors.NewIntrenalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
			logger.Error("error when scan user row  into users struct", err)
			return nil, errors.NewIntrenalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
