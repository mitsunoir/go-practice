package book

import (
	"context"
)

type Service struct {
	BookRepo  Repository
	Validator *Validator
}

func NewService(repo Repository, validator *Validator) *Service {
	return &Service{repo, validator}
}

func (s *Service) FindBook(ctx context.Context, id string) (*Book, error) {
	book, err := s.BookRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

type CreateBookRequest struct {
	Title   string
	Authors []string
}

func (s *Service) CreateBook(ctx context.Context, req *CreateBookRequest) (*Book, error) {
	if err := s.Validator.IsNotEmptyTitle(ctx, req); err != nil {
		return nil, err
	}
	if err := s.Validator.IsAuthorExists(ctx, req); err != nil {
		return nil, err
	}

	b, err := NewBook(req.Title, req.Authors)
	if err != nil {
		return nil, err
	}
	b, err = s.BookRepo.Create(ctx, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
