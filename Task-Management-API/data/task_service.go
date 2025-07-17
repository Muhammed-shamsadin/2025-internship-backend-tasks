package data

import (
	"2025-internship-backend-tasks/Task-Management-API/models"
	"errors"
)

// var tasks = []models.Task{}

type TaskService struct {
	tasks []models.Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: []models.Task{}, // In a real app, this would be a database connection
	}
}

// GetAllTasks
func (ts *TaskService) GetAllTasks() []models.Task {
	return ts.tasks
}

// GetTaskByID returns a task by its ID
func (ts *TaskService) GetTaskByID(id string) (*models.Task, error) {
	for _, task := range ts.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

//Update tasks
func (ts *TaskService) UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	for i, task := range ts.tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				ts.tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				ts.tasks[i].Description = updatedTask.Description
			}
			return &ts.tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}
