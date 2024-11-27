package task

import (
	"github.com/gin-gonic/gin"
)

const TASK_PATH string = "/task"

func CreateTaskRoutes(router *gin.Engine) {
	router.GET(TASK_PATH, getAllTasksHandler)
	router.POST(TASK_PATH, postNewTaskHandler)
	router.DELETE(TASK_PATH + "/:id", deleteTaskHandler)
	router.PATCH(TASK_PATH + "/:id", updateTaskHandler)
}