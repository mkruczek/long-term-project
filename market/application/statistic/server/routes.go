package server

import (
	"fmt"
	"market/market/application/statistic/handlers"
)

func (svr *Server) Routes() {

	mainGroup := svr.Group(fmt.Sprintf("%s", svr.UrlPath))

	// statistic part
	statsGroup := mainGroup.Group("/stats")
	statsGroup.POST("", handlers.Stats(svr.services.sts))
}
