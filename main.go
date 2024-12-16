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
		taskRoutes.POST("/tasks", controllers.CreateTask)
		taskRoutes.GET("/tasks", controllers.GetTask)
		taskRoutes.PUT("/tasks/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/tasks/:id", controllers.DeleteTask)

	}
	r.Run(":8080")

}
