package domain

import (
	"context"
	"time"
)

type getter interface {
	Get(ctx context.Context, id string) (Trade, error)
	GetAll(ctx context.Context) ([]Trade, error)
	GetRange(ctx context.Context, startTime, endTime time.Time) ([]Trade, error)
}

type Service struct {
	repo getter
}

func NewService(get getter) Service {
	return Service{
		repo: get,
	}
}

func (s Service) List(ctx context.Context) ([]Trade, error) {
	return s.repo.GetAll(ctx)
}

func (s Service) Get(ctx context.Context, id string) (Trade, error) {
	return s.repo.Get(ctx, id)
}

func (s Service) GetRange(ctx context.Context, startTime, endTime time.Time) ([]Trade, error) {
	return s.repo.GetRange(ctx, startTime, endTime)
}
