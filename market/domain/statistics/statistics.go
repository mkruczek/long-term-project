// Package statistics provides calculation of statistics for trades
package statistics

import (
	"market/market/domain"
	"math"
)

type Summary struct {
	// Profit is the sum of all profits in points
	Profit int `json:"profit"`
	// AverageProfit is the average profit rounded to the nearest integer
	// I chose to round to the nearest integer because at the end this is result in points, not pips
	AverageProfit int `json:"averageProfit"`
}

func Calculate(trades []domain.Trade) Summary {

	if len(trades) == 0 {
		return Summary{}
	}

	pro := profit(trades)
	averagePro := int(math.Round(float64(pro) / float64(len(trades))))

	return Summary{
		Profit:        pro,
		AverageProfit: averagePro,
	}
}

func profit(trades []domain.Trade) int {
	var result int
	for _, trade := range trades {
		result += trade.Profit
	}
	return result
}
