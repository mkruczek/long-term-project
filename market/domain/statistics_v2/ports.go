package statistics_v2

import "context"

type getTrades interface {
	GetTrades(ctx context.Context, filter Filter) ([]trade, error)
}
