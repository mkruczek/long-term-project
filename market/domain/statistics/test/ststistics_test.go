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

func TestCalculate_bestTrade_worstTrade(t *testing.T) {

	testCases := []struct {
		name        string
		trades      []domain.Trade
		bestProfit  int
		worstProfit int
	}{
		{name: "nil", trades: nil, bestProfit: 0, worstProfit: 0},
		{name: "empty", trades: []domain.Trade{}, bestProfit: 0, worstProfit: 0},
		{name: "one trade", trades: []domain.Trade{{Profit: 10}}, bestProfit: 10, worstProfit: 10},
		{name: "simple trades", trades: []domain.Trade{{Profit: -10}, {Profit: 20}, {Profit: 30}}, bestProfit: 30, worstProfit: -10},
		{name: "all trades are equal", trades: []domain.Trade{{Profit: 10}, {Profit: 10}, {Profit: 10}}, bestProfit: 10, worstProfit: 10},
		{name: "all are in plus", trades: []domain.Trade{{Profit: 10}, {Profit: 20}, {Profit: 30}}, bestProfit: 30, worstProfit: 10},
		{name: "all are in minus", trades: []domain.Trade{{Profit: -10}, {Profit: -20}, {Profit: -30}}, bestProfit: -10, worstProfit: -30},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			summary := statistics.Calculate(tc.trades)

			if summary.BestTrade.Profit != tc.bestProfit {
				t.Errorf("Expected best trade to be %v, got %v", tc.bestProfit, summary.BestTrade.Profit)
			}

			if summary.WorstTrade.Profit != tc.worstProfit {
				t.Errorf("Expected worst trade to be %v, got %v", tc.worstProfit, summary.WorstTrade.Profit)
			}
		})
	}
}
