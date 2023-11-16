package statistics_v2

import (
	"context"
	aclTrade "market/market/domain/trade"
	"market/market/libs/mongo"
)

type tradesRepository struct {
	db mongo.Provider
}

func (r *tradesRepository) GetTrades(ctx context.Context, filter Filter) ([]trade, error) {
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

func domainToTrade(domainTrades []aclTrade.Trade) []trade {
	trades := make([]trade, len(domainTrades))
	for i, t := range domainTrades {
		trades[i] = trade{
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
