package db

import (
	"context"
	"fmt"

	author "github.com/mitsunoir/go-practice/domain/author"
)

var authorDB = make(map[string]*author.Author)

type AuthorRepository struct{}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{}
}

func (r *AuthorRepository) FindByID(ctx context.Context, id string) (*author.Author, error) {
	if author, ok := authorDB[id]; !ok {
		return nil, fmt.Errorf("Author(id=%s) not found", id)
	} else {
		return author, nil
	}
}
func (r *AuthorRepository) Create(ctx context.Context, a *author.Author) (*author.Author, error) {
	authorDB[a.ID] = a
	return a, nil
}
