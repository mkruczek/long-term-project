package domain

import (
	"context"
	"market/market/infrastructure/log"
	"time"
)

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

func (s Service) Profit(ctx context.Context, startTime, endTime time.Time) (int, error) {

	//todo: get trades with time range
	list, err := s.List(ctx)
	if err != nil {
		return 0, err
	}

	var result int

	for _, trade := range list {
		if trade.OpenTime.After(startTime) && trade.CloseTime.Before(endTime) {
			log.Debugf(ctx, "trade: %s in range: %v - %v with profit: %d ", trade.ID, trade.OpenTime, trade.CloseTime, trade.Profit)
			result += trade.Profit
		}

	}

	return result, nil
}
