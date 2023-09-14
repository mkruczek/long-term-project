// Package statistics provides calculation of statistics for trades
package statistics

import (
	"market/market/domain"
	"math"
	"sync"
)

type Summary struct {
	// Profit is the sum of all profits in points
	Profit int `json:"profit"`
	// AverageProfit is the average profit rounded to the nearest integer
	// I chose to round to the nearest integer because at the end this is result in points, not pips
	AverageProfit int `json:"averageProfit"`
	// BestTrade is the trade with the highest profit
	BestTrade domain.Trade `json:"bestTrade"`
	// WorstTrade is the trade with the lowest profit
	WorstTrade domain.Trade `json:"worstTrade"`
}

type characteristic int

func (c characteristic) String() string {
	switch c {
	case best:
		return "best"
	case worst:
		return "worst"
	default:
		return "unknown"
	}
}

const (
	unknown characteristic = iota
	best
	worst
)

// tradesMap is a map of trades by their characteristics: worst, best, etc.
type tradesMap map[characteristic]domain.Trade

type tradeWithCharacteristic struct {
	trade          domain.Trade
	characteristic characteristic
}

// todo figure out how to do all inner calculations in generic way
func Calculate(trades []domain.Trade) Summary {

	if len(trades) == 0 {
		return Summary{}
	}

	pro := profit(trades)
	averagePro := int(math.Round(float64(pro) / float64(len(trades))))

	resultChan := make(chan tradeWithCharacteristic, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go bestTrade(wg, trades, resultChan)
	go worstTrade(wg, trades, resultChan)

	wg.Wait()
	close(resultChan)

	tradesMap := make(tradesMap)
	for result := range resultChan {
		tradesMap[result.characteristic] = result.trade
	}

	return Summary{
		Profit:        pro,
		AverageProfit: averagePro,
		BestTrade:     tradesMap[best],
		WorstTrade:    tradesMap[worst],
	}
}

func profit(trades []domain.Trade) int {
	var result int
	for _, trade := range trades {
		result += trade.Profit
	}
	return result
}

func bestTrade(wg *sync.WaitGroup, trades []domain.Trade, resultChan chan<- tradeWithCharacteristic) {
	defer wg.Done()
	var b domain.Trade
	for _, trade := range trades {
		if trade.Profit > b.Profit {
			b = trade
		}
	}
	resultChan <- tradeWithCharacteristic{trade: b, characteristic: best}

}

func worstTrade(wg *sync.WaitGroup, trades []domain.Trade, resultChan chan<- tradeWithCharacteristic) {
	defer wg.Done()
	var w domain.Trade
	for _, trade := range trades {
		if trade.Profit < w.Profit {
			w = trade
		}
	}
	resultChan <- tradeWithCharacteristic{trade: w, characteristic: worst}
}
