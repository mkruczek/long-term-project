package test

import (
	"github.com/google/go-cmp/cmp"
	"market/market/domain"
	"market/market/domain/statistics"
	"math"
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

func TestCalculate_stats_by_symbol(t *testing.T) {

	testCases := []struct {
		name     string
		trades   []domain.Trade
		expected statistics.Summary
	}{
		{name: "one pair with three trades", trades: []domain.Trade{
			{Symbol: "EURUSD", Profit: 10},
			{Symbol: "EURUSD", Profit: 20},
			{Symbol: "EURUSD", Profit: 30},
		},
			expected: statistics.Summary{
				BySymbol: map[string]statistics.BySymbol{
					"EURUSD": {Profit: 60, AverageProfit: 20, Amount: 3, PercentOfAll: 100},
				},
			},
		},
		{name: "two pairs with three trades", trades: []domain.Trade{
			{Symbol: "EURUSD", Profit: 10}, {Symbol: "EURUSD", Profit: 20}, {Symbol: "EURUSD", Profit: 30},
			{Symbol: "EURGBP", Profit: 10}, {Symbol: "EURGBP", Profit: 20}, {Symbol: "EURGBP", Profit: 30},
		},
			expected: statistics.Summary{
				BySymbol: map[string]statistics.BySymbol{
					"EURUSD": {Profit: 60, AverageProfit: 20, Amount: 3, PercentOfAll: 50},
					"EURGBP": {Profit: 60, AverageProfit: 20, Amount: 3, PercentOfAll: 50},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			summary := statistics.Calculate(tc.trades)
			if len(summary.BySymbol) != len(tc.expected.BySymbol) {
				t.Errorf("Expected %d symbols, got %d", len(tc.expected.BySymbol), len(summary.BySymbol))
			}
			for k, v := range summary.BySymbol {
				if !cmp.Equal(v, tc.expected.BySymbol[k]) {
					t.Errorf("Expected %v, got %v", tc.expected.BySymbol[k], v)
				}
			}
		})
	}
}

func Test_winLossRatio(t *testing.T) {
	testCases := []struct {
		name          string
		trades        []domain.Trade
		expectedRatio float64
	}{
		{name: "nil", trades: nil, expectedRatio: 0},
		{name: "empty", trades: []domain.Trade{}, expectedRatio: 0},
		{name: "one win trade", trades: []domain.Trade{{SimplifiedResult: domain.Win}}, expectedRatio: 1},
		{name: "one loss trade", trades: []domain.Trade{{SimplifiedResult: domain.Loss}}, expectedRatio: 0},
		{name: "one breakeven trade", trades: []domain.Trade{{SimplifiedResult: domain.BreakEven}}, expectedRatio: 0},
		{name: "one win and one loss trade", trades: []domain.Trade{{SimplifiedResult: domain.Win}, {SimplifiedResult: domain.Loss}}, expectedRatio: 0.5},
		{name: "one win and one loss and one breakeven trade", trades: []domain.Trade{{SimplifiedResult: domain.Win}, {SimplifiedResult: domain.Loss}, {SimplifiedResult: domain.BreakEven}}, expectedRatio: 0.5},
		{name: "two win and one loss trade", trades: []domain.Trade{{SimplifiedResult: domain.Win}, {SimplifiedResult: domain.Win}, {SimplifiedResult: domain.Loss}}, expectedRatio: 0.67},
		{name: "two win and one loss and one breakeven trade", trades: []domain.Trade{{SimplifiedResult: domain.Win}, {SimplifiedResult: domain.Win}, {SimplifiedResult: domain.Loss}, {SimplifiedResult: domain.BreakEven}}, expectedRatio: 0.67},
		{name: "two win and two loss trade", trades: []domain.Trade{{SimplifiedResult: domain.Win}, {SimplifiedResult: domain.Win}, {SimplifiedResult: domain.Loss}, {SimplifiedResult: domain.Loss}}, expectedRatio: 0.5},
		{name: "two win and two loss and one breakeven trade", trades: []domain.Trade{{SimplifiedResult: domain.Win}, {SimplifiedResult: domain.Win}, {SimplifiedResult: domain.Loss}, {SimplifiedResult: domain.Loss}, {SimplifiedResult: domain.BreakEven}}, expectedRatio: 0.5},
		{name: "one win and two loss trade", trades: []domain.Trade{{SimplifiedResult: domain.Win}, {SimplifiedResult: domain.Loss}, {SimplifiedResult: domain.Loss}}, expectedRatio: 0.33},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			summary := statistics.Calculate(tc.trades)
			if math.Round(summary.WinLossRatio*100)/100 != tc.expectedRatio {
				t.Errorf("Expected winLossRatio to be %v, got %v", tc.expectedRatio, summary.WinLossRatio)
			}
		})
	}
}
