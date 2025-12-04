package author

import "context"

type Service struct {
	AuthorRepo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) FindAuthor(ctx context.Context, id string) (*Author, error) {
	author, err := s.AuthorRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return author, nil
}

type CreateAuthorRequest struct {
	Name string
}

func (s *Service) CreateAuthor(ctx context.Context, req *CreateAuthorRequest) (*Author, error) {
	a, err := NewAuthor(req.Name)
	if err != nil {
		return nil, err
	}
	a, err = s.AuthorRepo.Create(ctx, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
