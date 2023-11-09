// Package statistics provides calculation of statistics for trades
package statistics

import (
	"context"
	"github.com/jinzhu/copier"
	"log/slog"
	"market/market/domain"
	"market/market/libs/log"
	"math"
	"sync"
)

type Summary struct {
	// Profit is the sum of all profits in points
	Profit int `json:"profit"`
	// AverageProfit is the average profit rounded to the nearest integer
	// I chose to round to the nearest integer because at the end this is result in points, not pips
	AverageProfit int `json:"averageProfit"`
	// WinLossRatio is the ratio of winning trades to losing trades
	// warning! break even trades are not taken into account
	WinLossRatio float64 `json:"winLossRatio"`
	// BestTrade is the trade with the highest profit
	BestTrade domain.Trade `json:"bestTrade"`
	// WorstTrade is the trade with the lowest profit
	WorstTrade domain.Trade `json:"worstTrade"`
	// BySymbol shows statistics for each symbol
	BySymbol map[string]BySymbol `json:"bySymbol"`
}

type BySymbol struct {
	// Profit is the sum of all profits in points
	Profit int `json:"profit"`
	// AverageProfit is the average profit rounded to the nearest integer
	// I chose to round to the nearest integer because at the end this is result in points, not pips
	AverageProfit int `json:"averageProfit"`
	// Amount is the number of trades
	Amount int `json:"amount"`
	// PercentOfAll is the percentage of all trades
	PercentOfAll int `json:"percentOfAll"`
}

func Calculate(trades []domain.Trade) Summary {

	if len(trades) == 0 {
		return Summary{}
	}

	const operations = 5

	resultChan := make(chan Summary, operations)
	wg := &sync.WaitGroup{}

	wg.Add(operations)
	go profit(wg, trades, resultChan)
	go bestTrade(wg, trades, resultChan)
	go worstTrade(wg, trades, resultChan)
	go calculateBySymbol(wg, trades, resultChan)
	go winLossRatio(wg, trades, resultChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

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
	slog.Debug("start calculating profit")
	defer wg.Done()
	var result int
	for _, trade := range trades {
		result += trade.Profit
	}
	resultChan <- Summary{Profit: result}
	slog.Debug("end calculating profit")
}

func bestTrade(wg *sync.WaitGroup, trades []domain.Trade, resultChan chan<- Summary) {
	slog.Debug("start calculating best trade")
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
	slog.Debug("end calculating best trade")
}

func worstTrade(wg *sync.WaitGroup, trades []domain.Trade, resultChan chan<- Summary) {
	slog.Debug("start calculating worst trade")
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
	slog.Debug("end calculating worst trade")
}

func calculateBySymbol(wg *sync.WaitGroup, trades []domain.Trade, resultChan chan<- Summary) {
	slog.Debug("start calculating by symbol")
	defer wg.Done()

	tmp := make(map[string][]domain.Trade, len(trades))
	for _, trade := range trades {
		tmp[trade.Symbol] = append(tmp[trade.Symbol], trade)
	}

	innerWg := &sync.WaitGroup{}
	innerWg.Add(len(tmp))

	type innerSummary struct {
		symbol   string
		bySymbol BySymbol
	}

	innerChan := make(chan innerSummary, len(tmp))

	for s, t := range tmp {
		go func(symbol string, trades []domain.Trade, allTrades int) {
			defer innerWg.Done()

			var profit int
			for _, trade := range trades {
				profit += trade.Profit
			}

			innerChan <- innerSummary{symbol: symbol, bySymbol: BySymbol{Profit: profit, AverageProfit: int(math.Round(float64(profit) / float64(len(trades)))), Amount: len(trades), PercentOfAll: int(math.Round(float64(len(trades)) / float64(allTrades) * 100))}}
		}(s, t, len(trades))
	}
	innerWg.Wait()
	close(innerChan)

	result := Summary{BySymbol: make(map[string]BySymbol, len(tmp))}
	for r := range innerChan {
		result.BySymbol[r.symbol] = r.bySymbol
	}

	resultChan <- result
	slog.Debug("end calculating by symbol")
}

func winLossRatio(wg *sync.WaitGroup, trades []domain.Trade, resultChan chan<- Summary) {
	slog.Debug("start calculating win loss ratio")
	defer wg.Done()
	var win, loss, breakeven float64
	for _, trade := range trades {
		switch trade.SimplifiedResult {
		case domain.Win:
			win++
		case domain.Loss:
			loss++
		case domain.BreakEven: //todo? how to handle breakeven? for now i just ignore it
			breakeven++
		}
	}

	s := Summary{}
	switch {
	case int(breakeven) == len(trades):
		s.WinLossRatio = 0
	case loss == 0:
		s.WinLossRatio = 1
	default:
		s.WinLossRatio = win / (win + loss)
	}
	resultChan <- s
	slog.Debug("end calculating win loss ratio")
}
