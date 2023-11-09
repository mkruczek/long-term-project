package server

import (
	"fmt"
	handlers2 "market/market/application/trade/handlers"
)

func (svr *Server) Routes() {

	mainGroup := svr.Group(fmt.Sprintf("%s", svr.UrlPath))

	uploadGroup := mainGroup.Group("/upload")
	uploadGroup.POST("/xtb", handlers2.XtbUpload(svr.services.xtb))

	tradeGroup := mainGroup.Group("/trades")
	tradeGroup.GET("", handlers2.ListTrades(svr.services.trades))
	tradeGroup.GET("/:id", handlers2.GetTrade(svr.services.trades))

	// statistic part
	statsGroup := mainGroup.Group("/stats")
	statsGroup.POST("", handlers2.Stats(svr.services.trades))

	//create image
	imageGroup := mainGroup.Group("/image")
	imageGroup.GET("", handlers2.CreateImage(svr.services.trades))
}
