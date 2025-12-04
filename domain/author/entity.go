package author

import (
	"errors"

	"github.com/google/uuid"
)

type Author struct {
	ID   string
	Name string
}

func NewAuthor(name string) (*Author, error) {
	if name == "" {
		return nil, errors.New("Author name must not be empty")
	}
	return &Author{
		ID:   uuid.NewString(),
		Name: name,
	}, nil
}
