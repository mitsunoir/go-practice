package book

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (*Book, error)
	Create(ctx context.Context, b *Book) (*Book, error)
}
