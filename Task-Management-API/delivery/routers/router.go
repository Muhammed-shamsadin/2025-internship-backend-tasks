package routers

import (
	"2025-internship-backend-tasks/Task-Management-API/delivery/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController, authMiddleware gin.HandlerFunc) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", userController.RegisterUser)
		auth.POST("/login", userController.LoginUser)
	}

	api := r.Group("/api")
	if authMiddleware != nil {
		api.Use(authMiddleware)
	}
	{
		api.GET("/tasks", taskController.GetAllTasks)
		api.GET("/tasks/:id", taskController.GetTaskByID)
		api.POST("/tasks", taskController.CreateTask)
		api.PUT("/tasks/:id", taskController.UpdateTask)
		api.DELETE("/tasks/:id", taskController.DeleteTask)
	}

	return r
}
