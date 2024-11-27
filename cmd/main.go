package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yonatansnir/go-task/internals/task"
)

func main() {
	r := gin.Default();
	
	task.CreateTaskRoutes(r);

	r.Run();
}