package usecase

import (
	"context"

	"github.com/mitsunoir/go-practice/domain/book"
)

type GetBookUseCase struct {
	bookRepo book.Repository
}

func NewGetBookUseCase(bookRepo book.Repository) *GetBookUseCase {
	return &GetBookUseCase{bookRepo}
}

func (uc *GetBookUseCase) Execute(ctx context.Context, id string) (*book.Book, error) {
	book, err := uc.bookRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
