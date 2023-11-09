package server

import (
	"market/market/application/service"
	"market/market/domain"
	"market/market/domain/dataProviders/xtb"
	"market/market/domain/statistics_v2"
)

// appServices - services for application
type appServices struct {
	trades service.Trades
	sts    statistics_v2.Service

	xtb domain.Provider[*xtb.CSV]
}

// NewAppServices - create new appServices
func NewAppServices(trd service.Trades, xtb domain.Provider[*xtb.CSV], sts statistics_v2.Service) appServices {
	return appServices{
		trades: trd,
		sts:    sts,
		xtb:    xtb,
	}
}
