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

func NewService(r getter) Service {
	return Service{repo: r}
}

func (s Service) List(ctx context.Context) ([]Trade, error) {
	return s.repo.GetAll(ctx)
}

func (s Service) Get(ctx context.Context, id string) (Trade, error) {
	return s.repo.Get(ctx, id)
}

func (s Service) Profit(ctx context.Context, startTime, endTime time.Time) (int, error) {

	list, err := s.repo.GetRange(ctx, startTime, endTime)
	if err != nil {
		return 0, err
	}

	var result int

	for _, trade := range list {
		result += trade.Profit
	}

	return result, nil
}
