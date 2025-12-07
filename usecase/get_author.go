package usecase

import (
	"context"

	"github.com/mitsunoir/go-practice/domain/author"
)

type GetAuthorUseCase struct {
	repo author.Repository
}

func NewGetAuthorUseCase(repo author.Repository) *GetAuthorUseCase {
	return &GetAuthorUseCase{repo}
}

func (uc *GetAuthorUseCase) Execute(ctx context.Context, id string) (*author.Author, error) {
	author, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return author, nil
}
