package app

import (
	"github.com/posol/bookstore_users_api/controllers/ping"
	"github.com/posol/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("ping", ping.Ping)

	router.POST("users", users.CreateUser)
	router.GET("users/:user_id", users.GetUser)
	//router.GET("users/search", users.SearchUser)
}
