package user

import "context"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(ctx context.Context, name string) error {
	u := &User{Name: name}
	return s.repo.Save(ctx, u)
}

func (s *Service) Get(ctx context.Context, id int64) (*User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
