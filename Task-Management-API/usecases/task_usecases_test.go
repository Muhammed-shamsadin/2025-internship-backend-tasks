package usecases

import (
	"context"
	"errors"
	"testing"
	"time"

	"2025-internship-backend-tasks/Task-Management-API/domain/task"
	"2025-internship-backend-tasks/Task-Management-API/domain/task/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskUsecase_CreateTask(t *testing.T) {
	mockTaskRepo := new(mocks.TaskRepository)
	taskUsecase := NewTaskUsecase(mockTaskRepo)

	t.Run("success", func(t *testing.T) {
		taskToCreate := &task.Task{
			Title:       "Test Task",
			Description: "This is a test task",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "Pending",
		}

		mockTaskRepo.On("Create", mock.Anything, mock.AnythingOfType("*task.Task")).Return(nil).Once()

		err := taskUsecase.CreateTask(context.Background(), taskToCreate)

		assert.NoError(t, err)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		taskToCreate := &task.Task{
			Title:       "Test Task",
			Description: "This is a test task",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "Pending",
		}
		dbError := errors.New("database error")

		mockTaskRepo.On("Create", mock.Anything, mock.AnythingOfType("*task.Task")).Return(dbError).Once()

		err := taskUsecase.CreateTask(context.Background(), taskToCreate)

		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		mockTaskRepo.AssertExpectations(t)
	})
}

func TestTaskUsecase_GetTaskByID(t *testing.T) {
	mockTaskRepo := new(mocks.TaskRepository)
	taskUsecase := NewTaskUsecase(mockTaskRepo)
	testID := primitive.NewObjectID()

	t.Run("success", func(t *testing.T) {
		expectedTask := &task.Task{
			ID:          testID.Hex(),
			Title:       "Test Task",
			Description: "This is a test task",
		}

		mockTaskRepo.On("FindByID", mock.Anything, testID.Hex()).Return(expectedTask, nil).Once()

		resultTask, err := taskUsecase.GetTaskByID(context.Background(), testID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, expectedTask, resultTask)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("not_found", func(t *testing.T) {
		notFoundError := errors.New("task not found")
		mockTaskRepo.On("FindByID", mock.Anything, testID.Hex()).Return(nil, notFoundError).Once()

		resultTask, err := taskUsecase.GetTaskByID(context.Background(), testID.Hex())

		assert.Error(t, err)
		assert.Nil(t, resultTask)
		assert.Equal(t, notFoundError, err)
		mockTaskRepo.AssertExpectations(t)
	})
}

func TestTaskUsecase_GetAllTasks(t *testing.T) {
	mockTaskRepo := new(mocks.TaskRepository)
	taskUsecase := NewTaskUsecase(mockTaskRepo)

	t.Run("success", func(t *testing.T) {
		expectedTasks := []task.Task{
			{ID: primitive.NewObjectID().Hex(), Title: "Task 1"},
			{ID: primitive.NewObjectID().Hex(), Title: "Task 2"},
		}

		mockTaskRepo.On("FindAll", mock.Anything).Return(expectedTasks, nil).Once()

		tasks, err := taskUsecase.GetAllTasks(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expectedTasks, tasks)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		dbError := errors.New("database error")
		mockTaskRepo.On("FindAll", mock.Anything).Return(nil, dbError).Once()

		tasks, err := taskUsecase.GetAllTasks(context.Background())

		assert.Error(t, err)
		assert.Nil(t, tasks)
		assert.Equal(t, dbError, err)
		mockTaskRepo.AssertExpectations(t)
	})
}

func TestTaskUsecase_UpdateTask(t *testing.T) {
	mockTaskRepo := new(mocks.TaskRepository)
	taskUsecase := NewTaskUsecase(mockTaskRepo)
	testID := primitive.NewObjectID()

	t.Run("success", func(t *testing.T) {
		taskToUpdate := &task.Task{
			Title: "Updated Task",
		}

		mockTaskRepo.On("Update", mock.Anything, testID.Hex(), mock.AnythingOfType("*task.Task")).Return(nil).Once()

		err := taskUsecase.UpdateTask(context.Background(), testID.Hex(), taskToUpdate)

		assert.NoError(t, err)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		taskToUpdate := &task.Task{
			Title: "Updated Task",
		}
		dbError := errors.New("update error")
		mockTaskRepo.On("Update", mock.Anything, testID.Hex(), mock.AnythingOfType("*task.Task")).Return(dbError).Once()

		err := taskUsecase.UpdateTask(context.Background(), testID.Hex(), taskToUpdate)

		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		mockTaskRepo.AssertExpectations(t)
	})
}

func TestTaskUsecase_DeleteTask(t *testing.T) {
	mockTaskRepo := new(mocks.TaskRepository)
	taskUsecase := NewTaskUsecase(mockTaskRepo)
	testID := primitive.NewObjectID()

	t.Run("success", func(t *testing.T) {
		mockTaskRepo.On("Delete", mock.Anything, testID.Hex()).Return(nil).Once()

		err := taskUsecase.DeleteTask(context.Background(), testID.Hex())

		assert.NoError(t, err)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		dbError := errors.New("delete error")
		mockTaskRepo.On("Delete", mock.Anything, testID.Hex()).Return(dbError).Once()

		err := taskUsecase.DeleteTask(context.Background(), testID.Hex())

		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		mockTaskRepo.AssertExpectations(t)
	})
}
