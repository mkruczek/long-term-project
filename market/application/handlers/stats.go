package handlers

import (
	"github.com/gin-gonic/gin"
	"market/market/application/service"
	"market/market/domain"
	"market/market/domain/statistics"
	"net/http"
	"time"
)

type statsQuery struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Symbol    string `json:"symbol"`
}

func Stats(trd service.Trades) gin.HandlerFunc {
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

		filter := domain.Filter{
			StartTime: startTime,
			EndTime:   endTime,
			Symbol:    sq.Symbol,
		}

		trades, err := trd.GetFiltered(ctx, filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sts := statistics.Calculate(trades)

		c.JSON(http.StatusOK, sts)
	}
}
