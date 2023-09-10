// Package statistics provides calculation of statistics for trades
package statistics

import "market/market/domain"

// Profit - POC for calculate profit, return only points value
func Profit(trades []domain.Trade) int {
	// Calculate profit for trades
	var profit int
	for _, trade := range trades {
		profit += trade.Profit
	}
	return profit
}
