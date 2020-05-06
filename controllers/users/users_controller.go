package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/posol/bookstore_users_api/domain/users"
	"github.com/posol/bookstore_users_api/services"
	"github.com/posol/bookstore_users_api/utils/errors"
)

func CreateUser(c *gin.Context) {
	fmt.Println("new request...")
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO: handle json err
		fmt.Println(err)
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	fmt.Println("user - ", user)

	result, restError := services.CreateUser(user)
	if restError != nil {
		fmt.Println(restError)
		// TODO: handle user creation error
		c.JSON(restError.Status, restError)
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
