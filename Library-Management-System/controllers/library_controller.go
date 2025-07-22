package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Muhammed-shamsadin/2025-internship-backend-tasks/Library-Management-System/models"
	"github.com/Muhammed-shamsadin/2025-internship-backend-tasks/Library-Management-System/services"
)


type LibraryController struct {
	LibManager services.LibraryManager
	Reader     *bufio.Reader
}

func NewLibraryController(libManager services.LibraryManager) *LibraryController {
	return &LibraryController{
		LibManager: libManager,
		Reader:     bufio.NewReader(os.Stdin),
	}
}

func (lc *LibraryController) readInt(prompt string) (int, error) {
	fmt.Print(prompt)
	input, _ := lc.Reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

func (lc *LibraryController) readString(prompt string) string {
	fmt.Print(prompt)
	input, _ := lc.Reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// AddBook prompts the user for book details and adds the book
func (lc *LibraryController) AddBook() {
	fmt.Println("--- Add New Book ---")
	id, err := lc.readInt("Enter Book ID: ")
	if err != nil {
		fmt.Println("Invalid ID format.")
		return
	}
	title := lc.readString("Enter Title: ")
	author := lc.readString("Enter Author: ")

	lc.LibManager.AddBook(models.Book{ID: id, Title: title, Author: author, Status: "Available"})
}

// RemoveBook prompts for book ID and removes the book
func (lc *LibraryController) RemoveBook() {
	fmt.Println("--- Remove Book ---")
	bookID, err := lc.readInt("Enter Book ID to remove: ")
	if err != nil {
		fmt.Println("Invalid ID format.")
		return
	}
	lc.LibManager.RemoveBook(bookID)
}

// BorrowBook prompts for book and member IDs and borrows the book
func (lc *LibraryController) BorrowBook() {
	fmt.Println("--- Borrow Book ---")
	bookID, err := lc.readInt("Enter Book ID to borrow: ")
	if err != nil {
		fmt.Println("Invalid Book ID format.")
		return
	}
	memberID, err := lc.readInt("Enter Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID format.")
		return
	}
	err = lc.LibManager.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}


// ReturnBook prompts for book and member IDs and returns the book
func (lc *LibraryController) ReturnBook() {
	fmt.Println("--- Return Book ---")
	bookID, err := lc.readInt("Enter Book ID to return: ")
	if err != nil {
		fmt.Println("Invalid Book ID format.")
		return
	}
	memberID, err := lc.readInt("Enter Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID format.")
		return
	}
	err = lc.LibManager.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

// ListAvailableBooks lists all available books
func (lc *LibraryController) ListAvailableBooks() {
	fmt.Println("--- Available Books ---")
	books := lc.LibManager.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No books available.")
		return
	}
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

// ListBorrowedBooks lists books borrowed by a member
func (lc *LibraryController) ListBorrowedBooks() {
	fmt.Println("--- Borrowed Books by Member ---")
	memberID, err := lc.readInt("Enter Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID format.")
		return
	}
	books := lc.LibManager.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Printf("No books borrowed by member %d.\n", memberID)
		return
	}
	fmt.Printf("Books borrowed by Member %d:\n", memberID)
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
