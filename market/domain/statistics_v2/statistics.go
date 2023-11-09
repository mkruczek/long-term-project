// Package statistics provides calculation of statistics for trades
package statistics_v2

import (
	"context"
	"github.com/jinzhu/copier"
	"log/slog"
	"market/market/libs/log"
	"math"
	"sync"
)

func calculate(trades []trade) Summary {

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

func profit(wg *sync.WaitGroup, trades []trade, resultChan chan<- Summary) {
	slog.Debug("start calculating profit")
	defer wg.Done()
	var result int
	for _, t := range trades {
		result += t.profit
	}
	resultChan <- Summary{Profit: result}
	slog.Debug("end calculating profit")
}

func bestTrade(wg *sync.WaitGroup, trades []trade, resultChan chan<- Summary) {
	slog.Debug("start calculating best trade")
	defer wg.Done()
	best := trade{
		profit: math.MinInt64,
	}

	for _, trade := range trades {
		if trade.profit > best.profit {
			best = trade
		}
	}
	resultChan <- Summary{BestTrade: best}
	slog.Debug("end calculating best trade")
}

func worstTrade(wg *sync.WaitGroup, trades []trade, resultChan chan<- Summary) {
	slog.Debug("start calculating worst trade")
	defer wg.Done()
	worst := trade{
		profit: math.MaxInt64,
	}
	for _, trade := range trades {
		if trade.profit < worst.profit {
			worst = trade
		}
	}
	resultChan <- Summary{WorstTrade: worst}
	slog.Debug("end calculating worst trade")
}

func calculateBySymbol(wg *sync.WaitGroup, allTrades []trade, resultChan chan<- Summary) {
	slog.Debug("start calculating by symbol")
	defer wg.Done()

	tradesBySymbol := make(map[string][]trade, len(allTrades))
	for _, t := range allTrades {
		tradesBySymbol[t.symbol] = append(tradesBySymbol[t.symbol], t)
	}

	innerWg := &sync.WaitGroup{}
	innerWg.Add(len(tradesBySymbol))

	type innerSummary struct {
		symbol   string
		bySymbol BySymbol
	}

	innerChan := make(chan innerSummary, len(tradesBySymbol))

	for symbol, trades := range tradesBySymbol {
		go func(symbol string, trades []trade, allTrades int) {
			defer innerWg.Done()

			var profit int
			for _, t := range trades {
				profit += t.profit
			}

			innerChan <- innerSummary{symbol: symbol, bySymbol: BySymbol{Profit: profit, AverageProfit: int(math.Round(float64(profit) / float64(len(trades)))), Amount: len(trades), PercentOfAll: int(math.Round(float64(len(trades)) / float64(allTrades) * 100))}}
		}(symbol, trades, len(allTrades))
	}
	innerWg.Wait()
	close(innerChan)

	result := Summary{BySymbol: make(map[string]BySymbol, len(tradesBySymbol))}
	for r := range innerChan {
		result.BySymbol[r.symbol] = r.bySymbol
	}

	resultChan <- result
	slog.Debug("end calculating by symbol")
}

func winLossRatio(wg *sync.WaitGroup, trades []trade, resultChan chan<- Summary) {
	slog.Debug("start calculating win loss ratio")
	defer wg.Done()
	var w, l, breakeven float64
	for _, t := range trades {
		switch t.simplifiedResult {
		case win:
			w++
		case loss:
			l++
		case breakEven: //todo? how to handle breakeven? for now i just ignore it
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
		s.WinLossRatio = w / (w + l)
	}
	resultChan <- s
	slog.Debug("end calculating win loss ratio")
}
