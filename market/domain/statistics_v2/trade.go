package statistics_v2

import (
	"market/market/libs/fxmoney"
	"time"
)

type tradeSide int8

const (
	UndefinedSide tradeSide = iota
	Buy
	Sell
)

const breakEvenDeviation = 20

type tradeResult int8

const (
	UndefinedResult tradeResult = iota
	Win
	Loss
	BreakEven
)

// Trade - internal representation of a Trade, for now entirely base on domain.Trade
type Trade struct {
	ID               string
	Symbol           string
	TradeSide        tradeSide
	OpenPrice        fxmoney.Price
	OpenTime         time.Time
	ClosePrice       fxmoney.Price
	CloseTime        time.Time
	Profit           int
	SimplifiedResult tradeResult
}
