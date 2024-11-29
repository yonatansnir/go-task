package user

import "github.com/gin-gonic/gin"

const USER_PATH string = "/user"

func CreateUserRoutes(router *gin.Engine) {
	router.GET(USER_PATH, getAllUsers)
	router.POST(USER_PATH, postNewUser)
}