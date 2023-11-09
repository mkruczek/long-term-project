package handlers

import (
	"github.com/gin-gonic/gin"
	"market/market/application/service"
)

func ListTrades(trades service.Trades) gin.HandlerFunc {
	return func(c *gin.Context) {
		trades, err := trades.List(c.Request.Context())
		if err != nil {
			c.String(500, err.Error())
			return
		}

		c.JSON(200, trades)
	}
}

func GetTrade(trades service.Trades) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		trade, err := trades.Get(c.Request.Context(), id)
		if err != nil {
			c.String(500, err.Error())
			return
		}

		c.JSON(200, trade)
	}
}
