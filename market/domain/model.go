package domain

import (
	"strings"
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
	OpenPrice Price `json:"price" bson:"price"`
	// OpenTime is the time the trade was opened.
	OpenTime time.Time `json:"openTime" bson:"openTime"`
	// ClosePrice is the price of the trade.
	ClosePrice Price `json:"closePrice" bson:"closePrice"`
	// CloseTime is the time the trade was closed.
	CloseTime time.Time `json:"closeTime" bson:"closeTime"`
	// Profit in points is the profit of the trade.
	// Points are the smallest unit of price change in the market.
	// 10 points = 1 pip
	// this value is more important than the profit in money for me.
	// if negative, it means a loss.
	Profit int `json:"profit" bson:"profit"`
	// ExternalID is the unique identifier of the trade in the broker system.
	ExternalID string `json:"externalId" bson:"externalId"`
}

// Price represents a price in the market.
type Price struct {
	// Value is the value of the price.
	Value float64 `json:"value" bson:"value"`
	// Currency is the currency of the price.
	// is the second part of the symbol under trade.
	// for example, if the symbol is EURUSD, the currency is USD.
	// for example, if the symbol is EURJPY, the currency is JPY.
	Currency string `json:"currency" bson:"currency"`
	// Coefficient is the coefficient of the price.
	// it is used to calculate the value of the price.
	// for example, if the value is 1.2345 and the coefficient is 10000,
	// the value of the price is 12345.
	// this is because there is difference between the value of the XXXUSD and XXXJPY.
	Coefficient int `json:"coefficient" bson:"coefficient"`
}

// CalculateProfit calculates the profit of the trade.
func (t *Trade) CalculateProfit() {
	op := int(t.OpenPrice.Value * float64(t.OpenPrice.Coefficient))
	cp := int(t.ClosePrice.Value * float64(t.ClosePrice.Coefficient))

	t.Profit = cp - op
}

func GetCoefficient(symbol string) int {
	if strings.Contains(symbol, "JPY") {
		return 1000
	}
	return 100000
}
