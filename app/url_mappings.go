package app

import (
	"github.com/posol/bookstore_users_api/controllers/ping"
	"github.com/posol/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("ping", ping.Ping)
	router.GET("api/internal/users/search", users.Search)
	router.POST("api/users", users.Create)
	router.GET("api/users/:user_id", users.Get)
	router.PUT("api/users/:user_id", users.Update)
	router.PATCH("api/users/:user_id", users.Update)
	router.DELETE("api/users/:user_id", users.Delete)
}
