package handlers

import (
	"github.com/gin-gonic/gin"
	"market/market/application/statistic/useCase"
	"market/market/domain/statistics_v2"
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

		c.JSON(http.StatusOK, sts)
	}
}
