package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Muhammed-shamsadin/2025-internship-backend-tasks/Library-Management-System/controllers"
	"github.com/Muhammed-shamsadin/2025-internship-backend-tasks/Library-Management-System/services"
)

func main() {
	libraryService := services.NewLibrary()
	controller := controllers.NewLibraryController(libraryService)
	reader := bufio.NewReader(os.Stdin)

	// Pre-populate with some members and books for easier testing
	// controller.LibManager.AddBook(models.Book{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Status: "Available"})
	// controller.LibManager.AddBook(models.Book{ID: 2, Title: "Effective Go", Author: "Google", Status: "Available"})
	// libraryService.Members[1] = models.Member{ID: 1, Name: "Alice"}
	// libraryService.Members[2] = models.Member{ID: 2, Name: "Bob"}

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			controller.AddBook()
		case "2":
			controller.RemoveBook()
		case "3":
			controller.BorrowBook()
		case "4":
			controller.ReturnBook()
		case "5":
			controller.ListAvailableBooks()
		case "6":
			controller.ListBorrowedBooks()
		case "7":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
