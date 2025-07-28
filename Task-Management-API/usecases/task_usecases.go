package usecases

import (
	"2025-internship-backend-tasks/Task-Management-API/domain/task"
	"context"
)

type TaskUsecase struct {
	taskRepo task.TaskRepository
}

func NewTaskUsecase(taskRepo task.TaskRepository) *TaskUsecase {
	return &TaskUsecase{taskRepo: taskRepo}
}

func (u *TaskUsecase) CreateTask(ctx context.Context, t *task.Task) error {
	return u.taskRepo.Create(ctx, t)
}

func (u *TaskUsecase) GetTaskByID(ctx context.Context, id string) (*task.Task, error) {
	return u.taskRepo.FindByID(ctx, id)
}

func (u *TaskUsecase) GetAllTasks(ctx context.Context) ([]task.Task, error) {
	return u.taskRepo.FindAll(ctx)
}

func (u *TaskUsecase) UpdateTask(ctx context.Context, id string, t *task.Task) error {
	return u.taskRepo.Update(ctx, id, t)
}

func (u *TaskUsecase) DeleteTask(ctx context.Context, id string) error {
	return u.taskRepo.Delete(ctx, id)
}
