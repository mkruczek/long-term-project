package service

import (
	"context"
	"market/market/domain"
	"market/market/domain/statistics_v2"
	"time"
)

type Trades interface {
	List(ctx context.Context) ([]domain.Trade, error)
	Get(ctx context.Context, id string) (domain.Trade, error)
	GetRange(ctx context.Context, startTime, endTime time.Time) ([]domain.Trade, error)
	GetFiltered(ctx context.Context, filter domain.Filter) ([]domain.Trade, error)
}

type Statistics interface {
	Calculate(ctx context.Context, filter statistics_v2.Filter) (statistics_v2.Summary, error)
}
