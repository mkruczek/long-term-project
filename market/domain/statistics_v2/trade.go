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

const breakEvenDeviation = 20

type tradeResult int8

const (
	undefinedResult tradeResult = iota
	win
	loss
	breakEven
)

// todo - this struct is public, for me this TechDebt, as long as i don't introduce database for statistics
// Trade - internal representation of a Trade, for now entirely base on domain.Trade
type Trade struct {
	id               string
	symbol           string
	tradeSide        tradeSide
	openPrice        fxmoney.Price
	openTime         time.Time
	closePrice       fxmoney.Price
	closeTime        time.Time
	profit           int
	simplifiedResult tradeResult
}
