package book

import "github.com/google/uuid"

type Book struct {
	ID      string
	Title   string
	Authors []string // author IDs
	// Status
}

func NewBook(title string, authors []string) (*Book, error) {
	return &Book{
		ID:      uuid.NewString(),
		Title:   title,
		Authors: authors,
	}, nil
}
