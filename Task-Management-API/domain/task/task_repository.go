package task

import (
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	FindByID(ctx context.Context, id string) (*Task, error)
	FindAll(ctx context.Context) ([]Task, error)
	Update(ctx context.Context, id string, task *Task) error
	Delete(ctx context.Context, id string) error
}
