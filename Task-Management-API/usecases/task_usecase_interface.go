package usecases

import (
	"2025-internship-backend-tasks/Task-Management-API/domain/task"
	"context"
)

type TaskUsecaseInterface interface {
	CreateTask(ctx context.Context, t *task.Task) error
	GetTaskByID(ctx context.Context, id string) (*task.Task, error)
	GetAllTasks(ctx context.Context) ([]task.Task, error)
	UpdateTask(ctx context.Context, id string, t *task.Task) error
	DeleteTask(ctx context.Context, id string) error
}
