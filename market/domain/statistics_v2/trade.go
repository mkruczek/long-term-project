package statistics_v2

import (
	"market/market/libs/fxmoney"
	"time"
)

type tradeSide int8

const (
	undefinedSide tradeSide = iota
	buy
	sell
)

func (ts tradeSide) String() string {
	switch ts {
	case buy:
		return "buy"
	case sell:
		return "sell"
	default:
		return "undefined"
	}
}

const breakEvenDeviation = 20

type tradeResult int8

const (
	undefinedResult tradeResult = iota
	win
	loss
	breakEven
)

func (tr tradeResult) String() string {
	switch tr {
	case win:
		return "win"
	case loss:
		return "loss"
	case breakEven:
		return "breakEven"
	default:
		return "undefined"
	}
}

// todo - this struct is public, for me this TechDebt, as long as i don't introduce database for statistics
// and also figure out how pass this data to client
// Trade - internal representation of a Trade, for now entirely base on domain.Trade
type Trade struct {
	Id               string
	Symbol           string
	TradeSide        tradeSide
	OpenPrice        fxmoney.Price
	OpenTime         time.Time
	ClosePrice       fxmoney.Price
	CloseTime        time.Time
	Profit           int
	SimplifiedResult tradeResult
}
