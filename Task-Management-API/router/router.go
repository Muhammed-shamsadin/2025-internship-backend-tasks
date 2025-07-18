package router

import (
	"2025-internship-backend-tasks/Task-Management-API/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	taskController := controllers.NewTaskController()

	// Task routes
	router.GET("/tasks", taskController.GetTasks)
	router.GET("/tasks/:id", taskController.GetTaskByID)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.POST("/tasks", taskController.AddTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)


	return router
}
