package service

import (
	"context"
	"market/market/domain"
	"time"
)

type Trades interface {
	List(ctx context.Context) ([]domain.Trade, error)
	Get(ctx context.Context, id string) (domain.Trade, error)
	GetRange(ctx context.Context, startTime, endTime time.Time) ([]domain.Trade, error)
}
