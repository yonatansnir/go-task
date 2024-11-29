package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yonatansnir/go-task/internals/task"
	"github.com/yonatansnir/go-task/internals/user"
)

func main() {
	r := gin.Default();
	
	task.CreateTaskRoutes(r);
	user.CreateUserRoutes(r);

	r.Run();
}