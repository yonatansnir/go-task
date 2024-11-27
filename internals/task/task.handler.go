package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/yonatansnir/go-task/pkg"
)

var tasks []TaskItem = []TaskItem{};

// Get all tasks
func getAllTasksHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tasks)
}

// Post new task
func postNewTaskHandler(ctx *gin.Context) {
	var taskRequestBody TaskRequestBody;

		if err := ctx.ShouldBindJSON(&taskRequestBody); err != nil {
			ctx.JSON(http.StatusNotAcceptable, gin.H{ "message": err.Error() })
			return
		}

		newTask := TaskItem{
			Title: taskRequestBody.Title,
			Description: taskRequestBody.Description,
			ID: utils.GenerateRandomID(),
			Done: false,
		}
		tasks = append(tasks, newTask);
		ctx.JSON(http.StatusCreated, newTask);
}


// Delete task by ID
func deleteTaskHandler (ctx *gin.Context) {
	id := ctx.Param("id");

	for i, task := range tasks {
		if (task.ID == id) {
			tasks = append(tasks[:i], tasks[i + 1:]...)
			ctx.JSON(http.StatusAccepted, gin.H{ "message": "id removed " + id })
			return
		}
	}

	ctx.JSON(http.StatusAccepted, gin.H{ "message": "Not found" })
}

// Update Task by ID
func updateTaskHandler (ctx *gin.Context) {
	id := ctx.Param("id");
	var taskRequestBody TaskRequestBody;

	if err := ctx.ShouldBindJSON(&taskRequestBody); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{ "message": err.Error() })
		return
	}

	updateTask(id, &taskRequestBody);
	ctx.JSON(http.StatusAccepted, gin.H{ "message": "Update successfully" })
}