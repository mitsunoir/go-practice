package usecase

import (
	"context"
	"errors"

	"github.com/mitsunoir/go-practice/domain/author"
	"github.com/mitsunoir/go-practice/domain/book"
)

type CreateBookUseCase struct {
	bookRepo   book.Repository
	authorRepo author.Repository
}

func NewCreateBookUseCase(bookRepo book.Repository, authorRepo author.Repository) *CreateBookUseCase {
	return &CreateBookUseCase{bookRepo, authorRepo}
}

type CreateBookRequest struct {
	Title   string
	Authors []string
}

func (uc *CreateBookUseCase) Execute(ctx context.Context, req *CreateBookRequest) (*book.Book, error) {
	if err := uc.isNotEmptyTitle(req); err != nil {
		return nil, err
	}
	if err := uc.isAuthorExists(ctx, req); err != nil {
		return nil, err
	}

	b, err := book.NewBook(req.Title, req.Authors)
	if err != nil {
		return nil, err
	}
	b, err = uc.bookRepo.Create(ctx, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (uc *CreateBookUseCase) isNotEmptyTitle(req *CreateBookRequest) error {
	if req.Title == "" {
		return errors.New("Book title must not be empty")
	}
	return nil
}

func (uc *CreateBookUseCase) isAuthorExists(ctx context.Context, req *CreateBookRequest) error {
	for _, authorID := range req.Authors {
		if _, err := uc.authorRepo.FindByID(ctx, authorID); err != nil {
			return err
		}
	}
	return nil
}
