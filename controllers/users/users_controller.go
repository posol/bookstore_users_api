package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/posol/bookstore_users_api/domain/users"
	"github.com/posol/bookstore_users_api/services"
)

func CreateUser(c *gin.Context) {
	fmt.Println("new request...")
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		// TODO: handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err)
		// TODO: handle json error
		return
	}
	fmt.Println("user - ", user)
	fmt.Println(string(bytes), "  ", err)
	
	result, err := services.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		// TODO: handle user creation error
		return
	}
	c.JSON(http.StatusCreated,result )
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement later")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement later")
}
