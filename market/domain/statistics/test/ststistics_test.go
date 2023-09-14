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
		{
			name: "trades with decimal up rounding",
			trades: []domain.Trade{
				{Profit: 100},
				{Profit: 100},
				{Profit: 105},
			},
			profit:  305,
			average: 102,
		},
		{
			name: "trades with decimal down rounding",
			trades: []domain.Trade{
				{Profit: 100},
				{Profit: 100},
				{Profit: 104},
			},
			profit:  304,
			average: 101,
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

func TestCalculate_BestTrade_WorstTrade(t *testing.T) {
	worst := domain.Trade{Profit: -100}
	best := domain.Trade{Profit: 100}

	trades := []domain.Trade{
		worst,
		best,
		{Profit: 10},
		{Profit: 20},
		{Profit: 30},
	}

	summary := statistics.Calculate(trades)

	if summary.BestTrade.Profit != best.Profit {
		t.Errorf("Expected best trade to be %v, got %v", best, summary.BestTrade)
	}

	if summary.WorstTrade.Profit != worst.Profit {
		t.Errorf("Expected worst trade to be %v, got %v", worst, summary.WorstTrade)
	}
}
