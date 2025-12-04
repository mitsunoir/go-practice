package db

import (
	"context"
	"fmt"

	book "github.com/mitsunoir/go-practice/domain/book"
)

var bookDB = make(map[string]*book.Book)

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) FindByID(ctx context.Context, id string) (*book.Book, error) {
	if book, ok := bookDB[id]; !ok {
		return nil, fmt.Errorf("Book(id=%s) not found", id)
	} else {
		return book, nil
	}
}
func (r *BookRepository) Create(ctx context.Context, b *book.Book) (*book.Book, error) {
	bookDB[b.ID] = b
	return b, nil
}
