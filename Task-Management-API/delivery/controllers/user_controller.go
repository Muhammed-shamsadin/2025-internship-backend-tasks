package controllers

import (
	"context"
	"net/http"
	"2025-internship-backend-tasks/Task-Management-API/domain/user"
	"2025-internship-backend-tasks/Task-Management-API/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase *usecases.UserUsecase
}

func NewUserController(usecase *usecases.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.usecase.RegisterUser(context.Background(), &u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (uc *UserController) LoginUser(c *gin.Context) {
	// Implement login logic here (call usecase, check password, return JWT)
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}
