package test

import (
	"market/market/domain"
	"market/market/domain/statistics"
	"testing"
)

func TestCalculate_Profit_AvrProfit(t *testing.T) {

	testCases := []struct {
		name    string
		trades  []domain.Trade
		profit  int
		average int
	}{
		{
			name:    "nil",
			trades:  nil,
			profit:  0,
			average: 0,
		},
		{
			name:    "empty",
			trades:  []domain.Trade{},
			profit:  0,
			average: 0,
		},
		{
			name: "simple trades",
			trades: []domain.Trade{
				{Profit: 10},
				{Profit: 20},
				{Profit: 30},
			},
			profit:  60,
			average: 20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			summary := statistics.Calculate(tc.trades)

			if summary.Profit != tc.profit {
				t.Errorf("Expected profit to be %d, got %d", tc.profit, summary.Profit)
			}

			if summary.AverageProfit != tc.average {
				t.Errorf("Expected average profit to be %d, got %d", tc.average, summary.AverageProfit)
			}
		})
	}
}
