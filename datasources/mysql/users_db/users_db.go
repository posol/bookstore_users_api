package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_users_user     = "mysql_users_user"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	user     = os.Getenv(mysql_users_user)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func int() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, password, host, schema)
	var err error

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	if err = Client.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("database successfully configured")
}
