// Package statistics provides calculation of statistics for trades
package statistics

import (
	"context"
	"github.com/jinzhu/copier"
	"market/market/domain"
	"market/market/infrastructure/log"
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

func Calculate(trades []domain.Trade) Summary {

	if len(trades) == 0 {
		return Summary{}
	}

	resultChan := make(chan Summary, 3)
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go profit(wg, trades, resultChan)
	go bestTrade(wg, trades, resultChan)
	go worstTrade(wg, trades, resultChan)

	wg.Wait()
	close(resultChan)

	var result Summary
	for r := range resultChan {
		err := copier.CopyWithOption(&result, &r, copier.Option{IgnoreEmpty: true})
		if err != nil {
			//todo? should i return error here?
			log.Errorf(context.Background(), "error collecting statistic for summary: %v", err)
			return Summary{}
		}
	}

	result.AverageProfit = int(math.Round(float64(result.Profit) / float64(len(trades))))

	return result
}

func profit(wg *sync.WaitGroup, trades []domain.Trade, resultChan chan<- Summary) {
	defer wg.Done()
	var result int
	for _, trade := range trades {
		result += trade.Profit
	}
	resultChan <- Summary{Profit: result}
}

func bestTrade(wg *sync.WaitGroup, trades []domain.Trade, resultChan chan<- Summary) {
	defer wg.Done()
	best := domain.Trade{
		Profit: math.MinInt64,
	}
	for _, trade := range trades {
		if trade.Profit > best.Profit {
			best = trade
		}
	}
	resultChan <- Summary{BestTrade: best}
}

func worstTrade(wg *sync.WaitGroup, trades []domain.Trade, resultChan chan<- Summary) {
	defer wg.Done()
	worst := domain.Trade{
		Profit: math.MaxInt64,
	}
	for _, trade := range trades {
		if trade.Profit < worst.Profit {
			worst = trade
		}
	}
	resultChan <- Summary{WorstTrade: worst}
}
