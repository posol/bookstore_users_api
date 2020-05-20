package mysql_errors

import (
	"github.com/go-sql-driver/mysql"
	"github.com/posol/bookstore_users_api/utils/errors"
	"strings"
)

const (
	errorsNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorsNoRows) {
			return errors.NewNotFoundError("no record matching with given id")
		}
		return errors.NewIntrenalServerError("error parsing mysql_errors db response")
	}

	switch sqlErr.Number {
	// duplicate record
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewIntrenalServerError("error processing request")
}
