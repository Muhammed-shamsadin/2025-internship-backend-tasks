package controllers

import (
	"2025-internship-backend-tasks/Task-Management-API/domain/task"
	"2025-internship-backend-tasks/Task-Management-API/usecases"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	usecase usecases.TaskUsecaseInterface
}

func NewTaskController(usecase usecases.TaskUsecaseInterface) *TaskController {
	return &TaskController{usecase: usecase}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var t task.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tc.usecase.CreateTask(context.Background(), &t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "task": t})
}

func (tc *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	t, err := tc.usecase.GetTaskByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": t})
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.usecase.GetAllTasks(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var t task.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tc.usecase.UpdateTask(context.Background(), id, &t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "task": t})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := tc.usecase.DeleteTask(context.Background(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
