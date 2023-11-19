package statistics_v2

import (
	"context"
	aclTrade "market/market/domain/trade"
	"market/market/libs/mongo"
)

// tradesRepository - i my mind this is  something which cane query the database and return the model of the domain
// todo? do i need import the Trade package model here?
// todo! fix dependency after introducing database for statistics
type TradesProvider struct {
	db *mongo.Provider
}

func NewTradesProvider(db *mongo.Provider) *TradesProvider {
	return &TradesProvider{db: db}
}

func (r *TradesProvider) GetTrades(ctx context.Context, filter Filter) ([]Trade, error) {
	var domainTrades []aclTrade.Trade
	var err error
	if filter.Symbol == "" {
		domainTrades, err = r.db.GetRange(ctx, filter.StartTime, filter.EndTime)
	} else {
		domainTrades, err = r.db.GetRangeAndSymbol(ctx, filter.StartTime, filter.EndTime, filter.Symbol)
	}

	if err != nil {
		return nil, err
	}

	return domainToTrade(domainTrades), nil
}

func domainToTrade(domainTrades []aclTrade.Trade) []Trade {
	trades := make([]Trade, len(domainTrades))
	for i, t := range domainTrades {
		trades[i] = Trade{
			id:               t.ID,
			symbol:           t.Symbol,
			tradeSide:        tradeSide(t.TradeSide),
			openPrice:        t.OpenPrice,
			openTime:         t.OpenTime,
			closePrice:       t.ClosePrice,
			closeTime:        t.CloseTime,
			profit:           t.Profit,
			simplifiedResult: calculateSimplifiedResult(t.Profit),
		}
	}
	return trades
}

func calculateSimplifiedResult(profit int) tradeResult {
	switch {
	case profit > breakEvenDeviation:
		return win
	case profit < -breakEvenDeviation:
		return loss
	default:
		return breakEven
	}
}
