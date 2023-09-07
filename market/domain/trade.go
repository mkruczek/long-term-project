package domain

import (
	"market/market/domain/fxmoney"
	"time"
)

type TradeSide int8

const (
	UndefinedSide TradeSide = iota
	Buy
	Sell
)

// Trade represents a trade in the market.
// This is the main domain model of this application.
// for now, we use the same this struct for both json and bson to check if it works.
type Trade struct {
	// ID is the unique identifier of the trade.
	ID string `json:"id" bson:"_id"`
	// Symbol is the symbol of the trade.
	Symbol string `json:"symbol" bson:"symbol"`
	// Type is the type of the trade.
	// Buy(Long) or Sell(Short)
	TradeSide TradeSide `json:"tradeSide" bson:"tradeSide"`
	// OpenPrice is the price of the trade.
	OpenPrice fxmoney.Price `json:"price" bson:"price"`
	// OpenTime is the time the trade was opened.
	OpenTime time.Time `json:"openTime" bson:"openTime"`
	// ClosePrice is the price of the trade.
	ClosePrice fxmoney.Price `json:"closePrice" bson:"closePrice"`
	// CloseTime is the time the trade was closed.
	CloseTime time.Time `json:"closeTime" bson:"closeTime"`
	// Profit in points is the profit of the trade.
	// Points are the smallest unit of price change in the market.
	// 10 points = 1 pip
	// this value is more important than the profit in fxmoney for me.
	// if negative, it means a loss.
	Profit int `json:"profit" bson:"profit"`
	// ExternalID is the unique identifier of the trade in the broker system.
	ExternalID string `json:"externalId" bson:"externalId"`
}

// CalculateProfit calculates the profit of the trade.
func (t *Trade) CalculateProfit() {

	var sign int
	switch t.TradeSide {
	case Buy:
		sign = 1
	case Sell:
		sign = -1
	}

	t.Profit = (t.ClosePrice.Subtract(t.OpenPrice).Amount) * sign
}
