package controllers

import (
	"2025-internship-backend-tasks/Task-Management-API/data"
	"2025-internship-backend-tasks/Task-Management-API/models"
	"2025-internship-backend-tasks/Task-Management-API/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	userService *data.UserService
}

func NewAuthController() *AuthController {
	return &AuthController{
		userService: data.NewUserService(),
	}
}



func (ac *AuthController) RegisterUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "Invalid input",
			"error": err.Error(),
		})
		return
	}

	// Check if username already exists
	exists, err := ac.userService.UsernameExists(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error checking user",
			"error":   err.Error(),
		})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Username already taken",
		})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password",
		})
		return
	}
	input.Password = string(hashedPassword)

	// Create the user
	if err := ac.userService.CreateUser(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

// LoginUser handles user login and JWT token generation
func (ac *AuthController) LoginUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	// Get user by username
	user, err := ac.userService.GetUserByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username or password",
		})
		return
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username or password",
		})
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
		})
		return
	}

	// Return token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}