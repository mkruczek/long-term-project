package server

import (
	"market/market/application/service"
	"market/market/domain"
	"market/market/domain/dataProviders/xtb"
)

// appServices - services for application
type appServices struct {
	trades service.Trades

	xtb domain.Provider[*xtb.CSV]
}

// NewAppServices - create new appServices
func NewAppServices(trd service.Trades, xtb domain.Provider[*xtb.CSV]) appServices {
	return appServices{
		trades: trd,
		xtb:    xtb,
	}
}
