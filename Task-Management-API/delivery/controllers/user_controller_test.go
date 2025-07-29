package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"2025-internship-backend-tasks/Task-Management-API/domain/user"
	"2025-internship-backend-tasks/Task-Management-API/usecases/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserController_RegisterUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(mocks.UserUsecaseInterface)
		userController := NewUserController(mockUsecase)

		router := gin.Default()
		router.POST("/register", userController.RegisterUser)

		newUser := &user.User{
			Username: "testuser",
			Password: "password123",
		}
		mockUsecase.On("RegisterUser", mock.Anything, mock.AnythingOfType("*user.User")).Return(nil).Once()

		body, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad_request", func(t *testing.T) {
		mockUsecase := new(mocks.UserUsecaseInterface)
		userController := NewUserController(mockUsecase)

		router := gin.Default()
		router.POST("/register", userController.RegisterUser)

		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUserController_LoginUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(mocks.UserUsecaseInterface)
	userController := NewUserController(mockUsecase)

	router := gin.Default()
	router.POST("/login", userController.LoginUser)

	t.Run("success", func(t *testing.T) {
		loginCreds := gin.H{
			"username": "testuser",
			"password": "password123",
		}
		userWithID := &user.User{
			ID:       "some-user-id",
			Username: "testuser",
		}

		mockUsecase.On("LoginUser", mock.Anything, "testuser", "password123").Return(userWithID, nil).Once()

		body, _ := json.Marshal(loginCreds)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "token")
		mockUsecase.AssertExpectations(t)
	})

	t.Run("invalid_credentials", func(t *testing.T) {
		loginCreds := gin.H{
			"username": "testuser",
			"password": "wrongpassword",
		}
		mockUsecase.On("LoginUser", mock.Anything, "testuser", "wrongpassword").Return(nil, errors.New("invalid credentials")).Once()

		body, _ := json.Marshal(loginCreds)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}
