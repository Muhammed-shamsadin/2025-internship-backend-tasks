package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"	
	"go.mongodb.org/mongo-driver/mongo/options"

	"2025-internship-backend-tasks/Task-Management-API/delivery/controllers"
	"2025-internship-backend-tasks/Task-Management-API/delivery/routers"
	"2025-internship-backend-tasks/Task-Management-API/repositories"
	"2025-internship-backend-tasks/Task-Management-API/usecases"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// MongoDB connection
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	db := client.Database("taskdb")

	// Repositories
	taskRepo := repositories.NewMongoTaskRepository(db.Collection("tasks"))
	userRepo := repositories.NewMongoUserRepository(db.Collection("users"))

	// Usecases
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo)

	// Controllers
	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	// Setup router (add your auth middleware as needed)
	r := routers.SetupRouter(taskController, userController)

	// Start server on port 8080
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
