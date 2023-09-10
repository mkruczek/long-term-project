package handlers

import (
	"github.com/gin-gonic/gin"
	"market/market/application/service"
	"net/http"
	"time"
)

type statsQuery struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type stats struct {
	Profit int `json:"profit"`
}

func Profit(sts service.Stats) gin.HandlerFunc {
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

		profit, err := sts.Profit(ctx, startTime, endTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := stats{
			Profit: profit,
		}

		c.JSON(http.StatusOK, response)
	}
}
