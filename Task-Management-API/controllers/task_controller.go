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
			"message": "Task not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task retrieved successfully",
		"task":    task,
	})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask models.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input, please provide a valid task",
			"error":   err.Error(),
		})
		return
	}

	task, err := tc.taskService.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to update task",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"task":    task,
	})
}

func (tc *TaskController) AddTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input, please provide a valid task",
			"error":   err.Error(),
		})
		return
	}

	task, err := tc.taskService.AddTask(newTask)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Failed to add task",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task added successfully",
		"task":    task,
	})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := tc.taskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to delete task",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
