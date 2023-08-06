package domain

import "time"

// Trade represents a trade in the market.
// This is the main domain model of this application.
// for now, we use the same this struct for both json and bson to check if it works.
type Trade struct {
	// ID is the unique identifier of the trade.
	ID string `json:"id" bson:"_id"`
	// Symbol is the symbol of the trade.
	Symbol string `json:"symbol" bson:"symbol"`
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
}

// Price represents a price in the market.
type Price struct {
	// Value is the value of the price.
	Value float64 `json:"value" bson:"value"`
	// Coefficient is the coefficient of the price.
	// it is used to calculate the value of the price.
	// for example, if the value is 1.2345 and the coefficient is 10000,
	// the value of the price is 12345.
	// this is because there is difference between the value of the XXXUSD and XXXJPY.
	Coefficient int `json:"coefficient" bson:"coefficient"`
}
