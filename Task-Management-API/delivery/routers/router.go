package routers

import (
	"2025-internship-backend-tasks/Task-Management-API/delivery/controllers"
	"2025-internship-backend-tasks/Task-Management-API/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", userController.RegisterUser)
		auth.POST("/login", userController.LoginUser)
	}

	api := r.Group("/api")
	api.Use(infrastructure.AuthMiddleware())
	{
		api.GET("/tasks", taskController.GetAllTasks)
		api.GET("/tasks/:id", taskController.GetTaskByID)
		api.POST("/tasks", taskController.CreateTask)
		api.PUT("/tasks/:id", taskController.UpdateTask)
		api.DELETE("/tasks/:id", taskController.DeleteTask)
	}

	return r
}
