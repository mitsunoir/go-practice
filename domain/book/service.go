package book

import (
	"context"
)

type Service struct {
	BookRepo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) FindBook(ctx context.Context, id string) (*Book, error) {
	book, err := s.BookRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *Service) CreateBook(ctx context.Context, b *Book) (*Book, error) {
	book, err := s.BookRepo.Create(ctx, b)
	if err != nil {
		return nil, err
	}
	return book, nil
}
