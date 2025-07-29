package usecases

import (
	"context"
	"errors"
	"testing"

	"2025-internship-backend-tasks/Task-Management-API/domain/user"
	"2025-internship-backend-tasks/Task-Management-API/domain/user/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_RegisterUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	userUsecase := NewUserUsecase(mockUserRepo)

	// Test case 1: Successful registration
	t.Run("success", func(t *testing.T) {
		newUser := &user.User{
			Username: "testuser",
			Password: "password123",
		}

		// Setup expectations
		mockUserRepo.On("UsernameExists", mock.Anything, "testuser").Return(false, nil).Once()
		mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*user.User")).Return(nil).Once()

		err := userUsecase.RegisterUser(context.Background(), newUser)

		// Assertions
		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})

	// Test case 2: Username already exists
	t.Run("username_exists", func(t *testing.T) {
		existingUser := &user.User{
			Username: "existinguser",
			Password: "password123",
		}

		// Setup expectations
		mockUserRepo.On("UsernameExists", mock.Anything, "existinguser").Return(true, nil).Once()

		err := userUsecase.RegisterUser(context.Background(), existingUser)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "username already exists", err.Error())
		mockUserRepo.AssertExpectations(t)
	})

	// Test case 3: Error checking if username exists
	t.Run("username_check_error", func(t *testing.T) {
		errorUser := &user.User{
			Username: "erroruser",
			Password: "password123",
		}
		dbError := errors.New("database error")

		// Setup expectations
		mockUserRepo.On("UsernameExists", mock.Anything, "erroruser").Return(false, dbError).Once()

		err := userUsecase.RegisterUser(context.Background(), errorUser)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_GetUserByUsername(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	userUsecase := NewUserUsecase(mockUserRepo)

	// Test case 1: Successful retrieval
	t.Run("success", func(t *testing.T) {
		expectedUser := &user.User{
			ID:       "some-id",
			Username: "testuser",
			Password: "password123",
		}

		// Setup expectations
		mockUserRepo.On("FindByUsername", mock.Anything, "testuser").Return(expectedUser, nil).Once()

		usr, err := userUsecase.GetUserByUsername(context.Background(), "testuser")

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, usr)
		mockUserRepo.AssertExpectations(t)
	})

	// Test case 2: User not found
	t.Run("not_found", func(t *testing.T) {
		notFoundError := errors.New("user not found")

		// Setup expectations
		mockUserRepo.On("FindByUsername", mock.Anything, "nonexistentuser").Return(nil, notFoundError).Once()

		usr, err := userUsecase.GetUserByUsername(context.Background(), "nonexistentuser")

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, usr)
		assert.Equal(t, notFoundError, err)
		mockUserRepo.AssertExpectations(t)
	})

	// Test case 3: Database error
	t.Run("database_error", func(t *testing.T) {
		dbError := errors.New("database error")

		// Setup expectations
		mockUserRepo.On("FindByUsername", mock.Anything, "db_error_user").Return(nil, dbError).Once()

		usr, err := userUsecase.GetUserByUsername(context.Background(), "db_error_user")

		//  
		assert.Error(t, err)
		assert.Nil(t, usr)
		assert.Equal(t, dbError, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_UsernameExists(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	userUsecase := NewUserUsecase(mockUserRepo)

	// Test case 1: Username exists
	t.Run("exists", func(t *testing.T) {
		// Setup expectations
		mockUserRepo.On("UsernameExists", mock.Anything, "existinguser").Return(true, nil).Once()

		exists, err := userUsecase.UsernameExists(context.Background(), "existinguser")

		// Assertions
		assert.NoError(t, err)
		assert.True(t, exists)
		mockUserRepo.AssertExpectations(t)
	})

	// Test case 2: Username does not exist
	t.Run("does_not_exist", func(t *testing.T) {
		// Setup expectations
		mockUserRepo.On("UsernameExists", mock.Anything, "newuser").Return(false, nil).Once()

		exists, err := userUsecase.UsernameExists(context.Background(), "newuser")

		// Assertions
		assert.NoError(t, err)
		assert.False(t, exists)
		mockUserRepo.AssertExpectations(t)
	})

	// Test case 3: Database error
	t.Run("database_error", func(t *testing.T) {
		dbError := errors.New("database error")

		// Setup expectations
		mockUserRepo.On("UsernameExists", mock.Anything, "db_error_user").Return(false, dbError).Once()

		exists, err := userUsecase.UsernameExists(context.Background(), "db_error_user")

		// Assertions
		assert.Error(t, err)
		assert.False(t, exists)
		assert.Equal(t, dbError, err)
		mockUserRepo.AssertExpectations(t)
	})
}
