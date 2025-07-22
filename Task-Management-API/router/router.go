package router

import (
	"2025-internship-backend-tasks/Task-Management-API/controllers"
	"2025-internship-backend-tasks/Task-Management-API/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	taskController := controllers.NewTaskController()
	authController := controllers.NewAuthController()
	middleware := middleware.NewAuthMiddleware()

	// Auth routes - Public routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.LoginUser)
		auth.POST("/register", authController.RegisterUser)
	}

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// Task routes
		api.GET("/tasks", taskController.GetTasks)
		api.GET("/tasks/:id", taskController.GetTaskByID)
		api.PUT("/tasks/:id", taskController.UpdateTask)
		api.POST("/tasks", taskController.AddTask)
		api.DELETE("/tasks/:id", taskController.DeleteTask)
	}


	return router
}
