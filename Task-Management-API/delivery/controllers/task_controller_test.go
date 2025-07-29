package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"2025-internship-backend-tasks/Task-Management-API/domain/task"
	"2025-internship-backend-tasks/Task-Management-API/usecases/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskController_CreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(mocks.TaskUsecaseInterface)
		taskController := NewTaskController(mockUsecase)

		router := gin.Default()
		router.POST("/tasks", taskController.CreateTask)

		newTask := &task.Task{
			Title:       "Test Task",
			Description: "This is a test task",
		}
		mockUsecase.On("CreateTask", mock.Anything, mock.AnythingOfType("*task.Task")).Return(nil).Once()

		body, _ := json.Marshal(newTask)
		req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}

func TestTaskController_GetTaskByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testID := primitive.NewObjectID()

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(mocks.TaskUsecaseInterface)
		taskController := NewTaskController(mockUsecase)

		router := gin.Default()
		router.GET("/tasks/:id", taskController.GetTaskByID)

		expectedTask := &task.Task{
			ID:          testID.Hex(),
			Title:       "Test Task",
			Description: "This is a test task",
			DueDate:     time.Now(),
			Status:      "Pending",
		}
		mockUsecase.On("GetTaskByID", mock.Anything, testID.Hex()).Return(expectedTask, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/tasks/"+testID.Hex(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}

func TestTaskController_GetAllTasks(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(mocks.TaskUsecaseInterface)
		taskController := NewTaskController(mockUsecase)

		router := gin.Default()
		router.GET("/tasks", taskController.GetAllTasks)

		expectedTasks := []task.Task{
			{ID: primitive.NewObjectID().Hex(), Title: "Task 1"},
			{ID: primitive.NewObjectID().Hex(), Title: "Task 2"},
		}
		mockUsecase.On("GetAllTasks", mock.Anything).Return(expectedTasks, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}

func TestTaskController_UpdateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testID := primitive.NewObjectID()

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(mocks.TaskUsecaseInterface)
		taskController := NewTaskController(mockUsecase)

		router := gin.Default()
		router.PUT("/tasks/:id", taskController.UpdateTask)

		updatedTask := &task.Task{
			Title: "Updated Task Title",
		}
		mockUsecase.On("UpdateTask", mock.Anything, testID.Hex(), mock.AnythingOfType("*task.Task")).Return(nil).Once()

		body, _ := json.Marshal(updatedTask)
		req, _ := http.NewRequest(http.MethodPut, "/tasks/"+testID.Hex(), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}

func TestTaskController_DeleteTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testID := primitive.NewObjectID()

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(mocks.TaskUsecaseInterface)
		taskController := NewTaskController(mockUsecase)

		router := gin.Default()
		router.DELETE("/tasks/:id", taskController.DeleteTask)

		mockUsecase.On("DeleteTask", mock.Anything, testID.Hex()).Return(nil).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+testID.Hex(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}
