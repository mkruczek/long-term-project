package useCase

import (
	"context"
	statistic "market/market/domain/statistics_v2"
)

type statistics interface {
	Calculate(ctx context.Context, filter statistic.Filter) (statistic.Summary, error)
}

// deprecated
// todo! remove this interface after introducing database for statistics
type getTrades interface {
	GetTrades(ctx context.Context, filter statistic.Filter) ([]statistic.Trade, error)
}

type StatisticGetter struct {
	Statistics statistics
	GetTrades  getTrades
}

func NewStatisticGetter(sts statistics, gt getTrades) StatisticGetter {
	return StatisticGetter{
		Statistics: sts,
		GetTrades:  gt,
	}
}
