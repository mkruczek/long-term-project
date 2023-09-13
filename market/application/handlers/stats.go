package handlers

import (
	"github.com/gin-gonic/gin"
	"market/market/application/service"
	"market/market/domain/statistics"
	"net/http"
	"time"
)

type statsQuery struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
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

		trades, err := trd.GetRange(ctx, startTime, endTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sts := statistics.Calculate(trades)

		c.JSON(http.StatusOK, sts)
	}
}
