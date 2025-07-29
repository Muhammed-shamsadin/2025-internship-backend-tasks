package controllers

import (
	"2025-internship-backend-tasks/Task-Management-API/domain/user"
	"2025-internship-backend-tasks/Task-Management-API/infrastructure"
	"2025-internship-backend-tasks/Task-Management-API/usecases"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase usecases.UserUsecaseInterface
}

func NewUserController(usecase usecases.UserUsecaseInterface) *UserController {
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
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	usr, err := uc.usecase.LoginUser(context.Background(), credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := infrastructure.GenerateJWT(usr.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
