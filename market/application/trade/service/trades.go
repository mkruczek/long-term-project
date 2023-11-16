package service

import (
	"context"
	"market/market/domain/statistics_v2"
	"market/market/domain/trade"
	"time"
)

type Trades interface {
	List(ctx context.Context) ([]trade.Trade, error)
	Get(ctx context.Context, id string) (trade.Trade, error)
	GetRange(ctx context.Context, startTime, endTime time.Time) ([]trade.Trade, error)
	GetFiltered(ctx context.Context, filter trade.Filter) ([]trade.Trade, error)
}

type Statistics interface {
	Calculate(ctx context.Context, filter statistics_v2.Filter) (statistics_v2.Summary, error)
}
