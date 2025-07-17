package main

import (
	"2025-internship-backend-tasks/Task-Management-API/router"
	"log"
)

func main() {
	// Setup router
	r := router.SetupRouter()

	// Start server on port 8080
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
