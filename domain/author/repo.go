package author

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (*Author, error)
	Create(ctx context.Context, a *Author) (*Author, error)
}
