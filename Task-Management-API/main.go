package main

import (
	"2025-internship-backend-tasks/Task-Management-API/data"
	"2025-internship-backend-tasks/Task-Management-API/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	//add mongo connection
	if err := data.ConnectMongoDB(); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	} // Setup router
	r := router.SetupRouter()

	// Start server on port 8080
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
