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

// trade - internal representation of a trade, for now entirely base on domain.Trade
type trade struct {
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
