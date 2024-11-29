package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users []User = []User{}

func getAllUsers (c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func postNewUser(c *gin.Context) {
	var newUser User;

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
		return;
	}

	if err := validateUserRequest(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() });
		return
	}

	users = append(users, newUser);

	c.JSON(http.StatusCreated, gin.H{ "message": "User created succesfully" })
}