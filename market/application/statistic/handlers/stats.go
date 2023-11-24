package handlers

import (
	"github.com/gin-gonic/gin"
	"market/market/application/statistic/useCase"
	"market/market/domain/statistics_v2"
	"market/market/libs/fxmoney"
	"net/http"
	"time"
)

type statsQuery struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Symbol    string `json:"symbol"`
}

func Stats(sg useCase.StatisticGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var sq statsQuery
		if err := c.ShouldBind(&sq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		startTime, err := time.Parse("2006-01-02T15:04:05Z", sq.StartTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		endTime, err := time.Parse("2006-01-02T15:04:05Z", sq.EndTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filter := statistics_v2.Filter{
			StartTime: startTime,
			EndTime:   endTime,
			Symbol:    sq.Symbol,
		}

		sts, err := sg.Statistics.Calculate(ctx, filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, domainModelToDto(sts))
	}
}

func domainModelToDto(d statistics_v2.Summary) summaryDto {
	dto := summaryDto{
		Profit:        d.Profit,
		AverageProfit: d.AverageProfit,
		WinLossRatio:  d.WinLossRatio,
		BestTrade:     tradeToDto(d.BestTrade),
		WorstTrade:    tradeToDto(d.WorstTrade),
		BySymbol:      make(map[string]bySymbol),
	}
	for k, v := range d.BySymbol {
		dto.BySymbol[k] = bySymbol{
			Profit:        v.Profit,
			AverageProfit: v.AverageProfit,
			Amount:        v.Amount,
			PercentOfAll:  v.PercentOfAll,
		}
	}
	return dto
}

func tradeToDto(t statistics_v2.Trade) trade {
	return trade{
		Id:               t.Id,
		Symbol:           t.Symbol,
		TradeSide:        t.TradeSide.String(),
		OpenPrice:        t.OpenPrice,
		OpenTime:         t.OpenTime,
		ClosePrice:       t.ClosePrice,
		CloseTime:        t.CloseTime,
		Profit:           t.Profit,
		SimplifiedResult: t.SimplifiedResult.String(),
	}
}

type summaryDto struct {
	// Profit is the sum of all profits in points
	Profit int `json:"profit"`
	// AverageProfit is the average profit rounded to the nearest integer
	// I chose to round to the nearest integer because at the end this is result in points, not pips
	AverageProfit int `json:"averageProfit"`
	// WinLossRatio is the ratio of winning trades to losing trades
	// warning! break even trades are not taken into account
	WinLossRatio float64 `json:"winLossRatio"`
	// BestTrade is the Trade with the highest profit
	BestTrade trade `json:"bestTrade"`
	// WorstTrade is the Trade with the lowest profit
	WorstTrade trade `json:"worstTrade"`
	// BySymbol shows statistics for each symbol
	BySymbol map[string]bySymbol `json:"bySymbol"`
}

type bySymbol struct {
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

type trade struct {
	Id               string        `json:"id"`
	Symbol           string        `json:"symbol"`
	TradeSide        string        `json:"tradeSide"`
	OpenPrice        fxmoney.Price `json:"openPrice"`
	OpenTime         time.Time     `json:"openTime"`
	ClosePrice       fxmoney.Price `json:"closePrice"`
	CloseTime        time.Time     `json:"closeTime"`
	Profit           int           `json:"profit"`
	SimplifiedResult string        `json:"simplifiedResult"`
}
