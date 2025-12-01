package book

import (
	"errors"

	"github.com/google/uuid"
)

type Book struct {
	ID      string
	Title   string
	Authors []string // author IDs
	// Status
}

func NewBook(title string, authors []string) (*Book, error) {
	if title == "" {
		return nil, errors.New("Book title must not be empty")
	}
	return &Book{
		ID:      uuid.NewString(),
		Title:   title,
		Authors: authors,
	}, nil
}
