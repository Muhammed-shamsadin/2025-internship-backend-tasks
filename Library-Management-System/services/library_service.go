package services

import (
	"fmt"

	"example.com/librarymanagement/models"
)

// LibraryManager Interface
type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

// AddBook adds a new book to the library
func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
	fmt.Printf("Book '%s' added successfully.\n", book.Title)
}

// RemoveBook removes a book from the library by its ID
func (l *Library) RemoveBook(bookID int) {
	book, exists := l.Books[bookID]
	if !exists {
		fmt.Printf("Book with ID %d not found.\n", bookID)
		return
	}
	delete(l.Books, bookID)
	fmt.Printf("Book '%s' removed successfully.\n", book.Title)
}

// BorrowBook allows a member to borrow a book if it is available
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	if !bookExists {
		return fmt.Errorf("book with ID %d not found", bookID)
	}
	if book.Status == "Borrowed" {
		return fmt.Errorf("book '%s' is already borrowed", book.Title)
	}

	// Check if member that is to borrow exist or not
	member, memberExists := l.Members[memberID]
	
	// if Memeber doesnt exist I am trying to register the memeber here
	if !memberExists {
		member = models.Member{ID: memberID, Name: fmt.Sprintf("Member %d", memberID), BorrowedBooks: []models.Book{}}
		l.Members[memberID] = member
		fmt.Printf("New member with ID %d created.\n", memberID)
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member

	fmt.Printf("Book '%s' borrowed by member %d successfully.\n", book.Title, memberID)
	return nil
}

// ReturnBook allows a member to return a borrowed book
func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	if !bookExists {
		return fmt.Errorf("book with ID %d not found", bookID)
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return fmt.Errorf("member with ID %d not found", memberID)
	}

	foundBookInMember := false
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			foundBookInMember = true
			break
		}
	}

	if !foundBookInMember {
		return fmt.Errorf("book '%s' was not borrowed by member %d", book.Title, memberID)
	}

	book.Status = "Available"
	l.Books[bookID] = book
	l.Members[memberID] = member

	fmt.Printf("Book '%s' returned by member %d successfully.\n", book.Title, memberID)
	return nil
}

// ListAvailableBooks lists all available books in the library
func (l *Library) ListAvailableBooks() []models.Book {
	availableBooks := []models.Book{}
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

// ListBorrowedBooks lists all books borrowed by a specific member
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		fmt.Printf("Member with ID %d not found.\n", memberID)
		return []models.Book{}
	}
	return member.BorrowedBooks
}
