package controllers

import (
	"2025-internship-backend-tasks/Task-Management-API/data"
	"2025-internship-backend-tasks/Task-Management-API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


// explain why we are using TaskController here ??
// The TaskController is responsible for handling HTTP requests related to tasks.
// It acts as an intermediary between the HTTP layer (Gin) and the service layer (TaskService).
// This separation of concerns makes the code more modular and easier to maintain.

type TaskController struct {
	taskService *data.TaskService
}


// explain the need of NewTaskController and what it does ?

func NewTaskController() *TaskController {
	return &TaskController{
		taskService: data.NewTaskService(),
	}
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks := tc.taskService.GetAllTasks()
	c.JSON(http.StatusOK, gin.H{
		"message": "List of all tasks",
		"tasks":   tasks,
	})
}

func (tc *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	task, err := tc.taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask models.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input, please provide a valid task",
		})
		return
	}

	task, err := tc.taskService.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	c.JSON(http.StatusOK, task)
}
