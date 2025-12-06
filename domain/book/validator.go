package book

import (
	"context"
	"errors"

	"github.com/mitsunoir/go-practice/domain/author"
)

type Validator struct {
	authorRepo author.Repository
}

func NewValidator(repo author.Repository) *Validator {
	return &Validator{repo}
}

func (v *Validator) IsNotEmptyTitle(ctx context.Context, req *CreateBookRequest) error {
	if req.Title == "" {
		return errors.New("Book title must not be empty")
	}
	return nil
}
func (v *Validator) IsAuthorExists(ctx context.Context, req *CreateBookRequest) error {
	for _, authorID := range req.Authors {
		if _, err := v.authorRepo.FindByID(ctx, authorID); err != nil {
			return err
		}
	}
	return nil
}
