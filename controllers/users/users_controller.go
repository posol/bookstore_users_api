package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/posol/bookstore_users_api/domain/users"
	"github.com/posol/bookstore_users_api/services"
)

func CreateUser(c *gin.Context) {
	fmt.Println("new request...")
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO: handle json err
		fmt.Println(err)
		return
	}
	fmt.Println("user - ", user)

	result, err := services.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		// TODO: handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement later")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement later")
}
