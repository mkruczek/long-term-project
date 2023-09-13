package server

import (
	"fmt"
	"market/market/application/handlers"
)

func (svr *Server) Routes() {

	mainGroup := svr.Group(fmt.Sprintf("%s", svr.UrlPath))

	uploadGroup := mainGroup.Group("/upload")
	uploadGroup.POST("/xtb", handlers.XtbUpload(svr.services.xtb))

	tradeGroup := mainGroup.Group("/trades")
	tradeGroup.GET("", handlers.ListTrades(svr.services.trades))
	tradeGroup.GET("/:id", handlers.GetTrade(svr.services.trades))

	// statistic part
	statsGroup := mainGroup.Group("/stats")
	statsGroup.POST("", handlers.Stats(svr.services.trades))
}
