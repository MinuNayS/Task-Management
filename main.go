package main

import (
	"task-manager/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.connect()
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.POST("/", controllers.CreateTask)
	}
	r.Run(":8080")

}
