package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/posol/bookstore_users_api/domain/users"
	"github.com/posol/bookstore_users_api/services"
	"github.com/posol/bookstore_users_api/utils/errors"
)

func getUserId(userIdParam string) (int64, *errors.RestError) {
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
	fmt.Println("new request...")
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	fmt.Println("user - ", user)

	result, restError := services.CreateUser(user)
	if restError != nil {
		c.JSON(restError.Status, restError)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	result, restError := services.GetUser(userId)
	if restError != nil {
		fmt.Println(restError)
		c.JSON(restError.Status, restError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	fmt.Println("new request...")
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	fmt.Println("user - ", user)

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, restError := services.UpdateUser(isPartial, user)
	if restError != nil {
		c.JSON(restError.Status, restError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement later")
}
