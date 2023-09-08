package service

import (
	"market/market/domain"
)

type Stats interface {
	Profit(trades []domain.Trade) (float64, error)
}
