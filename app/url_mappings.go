package app

import (
	"github.com/posol/bookstore_users_api/controllers/ping"
	"github.com/posol/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("ping", ping.Ping)

	router.POST("api/users", users.CreateUser)
	router.GET("api/users/:user_id", users.GetUser)
	router.PUT("api/users/:user_id", users.UpdateUser)
	router.PATCH("api/users/:user_id", users.UpdateUser)
	//router.GET("api/users/search", users.SearchUser)
}
