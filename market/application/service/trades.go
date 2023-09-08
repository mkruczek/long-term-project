package service

import (
	"context"
	"market/market/domain"
)

type Trades interface {
	List(ctx context.Context) ([]domain.Trade, error)
	Get(ctx context.Context, id string) (domain.Trade, error)
}
