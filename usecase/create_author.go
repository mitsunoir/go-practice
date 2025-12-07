package usecase

import (
	"context"

	"github.com/mitsunoir/go-practice/domain/author"
)

type CreateAuthorUseCase struct {
	repo author.Repository
}

func NewCreateAuthorUseCase(repo author.Repository) *CreateAuthorUseCase {
	return &CreateAuthorUseCase{repo}
}

type CreateAuthorRequest struct {
	Name string
}

func (uc *CreateAuthorUseCase) Execute(ctx context.Context, req *CreateAuthorRequest) (*author.Author, error) {
	a, err := author.NewAuthor(req.Name)
	if err != nil {
		return nil, err
	}
	a, err = uc.repo.Create(ctx, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
