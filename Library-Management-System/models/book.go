package models

type Book struct {
	ID int
	Title string
	Author string
	// Status => "Available" or "Borrowed"
	Status string
}
