package domain

import "context"

type getter interface {
	List(ctx context.Context) ([]Trade, error)
	Get(ctx context.Context, id string) (Trade, error)
}

type Service struct {
	repo getter
}

func NewService(r getter) Service {
	return Service{repo: r}
}

func (s Service) List(ctx context.Context) ([]Trade, error) {
	return s.repo.List(ctx)
}

func (s Service) Get(ctx context.Context, id string) (Trade, error) {
	return s.repo.Get(ctx, id)
}
