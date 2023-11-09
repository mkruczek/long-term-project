package statistics_v2

import (
	"github.com/google/go-cmp/cmp"
	"math"
	"testing"
)

func TestCalculate_Profit_AvrProfit(t *testing.T) {
	testCases := []struct {
		name    string
		trades  []trade
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
			trades:  []trade{},
			profit:  0,
			average: 0,
		},
		{
			name: "simple trades",
			trades: []trade{
				{profit: 10},
				{profit: 20},
				{profit: 30},
			},
			profit:  60,
			average: 20,
		},
		{
			name: "trades with decimal up rounding",
			trades: []trade{
				{profit: 100},
				{profit: 100},
				{profit: 105},
			},
			profit:  305,
			average: 102,
		},
		{
			name: "trades with decimal down rounding",
			trades: []trade{
				{profit: 100},
				{profit: 100},
				{profit: 104},
			},
			profit:  304,
			average: 101,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			summary := calculate(tc.trades)

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
		trades      []trade
		bestProfit  int
		worstProfit int
	}{
		{name: "nil", trades: nil, bestProfit: 0, worstProfit: 0},
		{name: "empty", trades: []trade{}, bestProfit: 0, worstProfit: 0},
		{name: "one trade", trades: []trade{{profit: 10}}, bestProfit: 10, worstProfit: 10},
		{name: "simple trades", trades: []trade{{profit: -10}, {profit: 20}, {profit: 30}}, bestProfit: 30, worstProfit: -10},
		{name: "all trades are equal", trades: []trade{{profit: 10}, {profit: 10}, {profit: 10}}, bestProfit: 10, worstProfit: 10},
		{name: "all are in plus", trades: []trade{{profit: 10}, {profit: 20}, {profit: 30}}, bestProfit: 30, worstProfit: 10},
		{name: "all are in minus", trades: []trade{{profit: -10}, {profit: -20}, {profit: -30}}, bestProfit: -10, worstProfit: -30},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			summary := calculate(tc.trades)

			if summary.BestTrade.profit != tc.bestProfit {
				t.Errorf("Expected best trade to be %v, got %v", tc.bestProfit, summary.BestTrade.profit)
			}

			if summary.WorstTrade.profit != tc.worstProfit {
				t.Errorf("Expected worst trade to be %v, got %v", tc.worstProfit, summary.WorstTrade.profit)
			}
		})
	}
}

func TestCalculate_stats_by_symbol(t *testing.T) {

	testCases := []struct {
		name     string
		trades   []trade
		expected Summary
	}{
		{name: "one pair with three trades", trades: []trade{
			{symbol: "EURUSD", profit: 10},
			{symbol: "EURUSD", profit: 20},
			{symbol: "EURUSD", profit: 30},
		},
			expected: Summary{
				BySymbol: map[string]BySymbol{
					"EURUSD": {Profit: 60, AverageProfit: 20, Amount: 3, PercentOfAll: 100},
				},
			},
		},
		{name: "two pairs with three trades", trades: []trade{
			{symbol: "EURUSD", profit: 10}, {symbol: "EURUSD", profit: 20}, {symbol: "EURUSD", profit: 30},
			{symbol: "EURGBP", profit: 10}, {symbol: "EURGBP", profit: 20}, {symbol: "EURGBP", profit: 30},
		},
			expected: Summary{
				BySymbol: map[string]BySymbol{
					"EURUSD": {Profit: 60, AverageProfit: 20, Amount: 3, PercentOfAll: 50},
					"EURGBP": {Profit: 60, AverageProfit: 20, Amount: 3, PercentOfAll: 50},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			summary := calculate(tc.trades)
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
		trades        []trade
		expectedRatio float64
	}{
		{name: "nil", trades: nil, expectedRatio: 0},
		{name: "empty", trades: []trade{}, expectedRatio: 0},
		{name: "one win trade", trades: []trade{{simplifiedResult: win}}, expectedRatio: 1},
		{name: "one loss trade", trades: []trade{{simplifiedResult: loss}}, expectedRatio: 0},
		{name: "one breakeven trade", trades: []trade{{simplifiedResult: breakEven}}, expectedRatio: 0},
		{name: "one win and one loss trade", trades: []trade{{simplifiedResult: win}, {simplifiedResult: loss}}, expectedRatio: 0.5},
		{name: "one win and one loss and one breakeven trade", trades: []trade{{simplifiedResult: win}, {simplifiedResult: loss}, {simplifiedResult: breakEven}}, expectedRatio: 0.5},
		{name: "two win and one loss trade", trades: []trade{{simplifiedResult: win}, {simplifiedResult: win}, {simplifiedResult: loss}}, expectedRatio: 0.67},
		{name: "two win and one loss and one breakeven trade", trades: []trade{{simplifiedResult: win}, {simplifiedResult: win}, {simplifiedResult: loss}, {simplifiedResult: breakEven}}, expectedRatio: 0.67},
		{name: "two win and two loss trade", trades: []trade{{simplifiedResult: win}, {simplifiedResult: win}, {simplifiedResult: loss}, {simplifiedResult: loss}}, expectedRatio: 0.5},
		{name: "two win and two loss and one breakeven trade", trades: []trade{{simplifiedResult: win}, {simplifiedResult: win}, {simplifiedResult: loss}, {simplifiedResult: loss}, {simplifiedResult: breakEven}}, expectedRatio: 0.5},
		{name: "one win and two loss trade", trades: []trade{{simplifiedResult: win}, {simplifiedResult: loss}, {simplifiedResult: loss}}, expectedRatio: 0.33},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			summary := calculate(tc.trades)
			if math.Round(summary.WinLossRatio*100)/100 != tc.expectedRatio {
				t.Errorf("Expected winLossRatio to be %v, got %v", tc.expectedRatio, summary.WinLossRatio)
			}
		})
	}
}
