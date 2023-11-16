package server

import (
	"market/market/application/trade/service"
	"market/market/domain/dataProviders/xtb"
	"market/market/domain/trade"
)

// appServices - services for application
type appServices struct {
	trades service.Trades

	xtb trade.Provider[*xtb.CSV]
}

// NewAppServices - create new appServices
func NewAppServices(trd service.Trades, xtb trade.Provider[*xtb.CSV]) appServices {
	return appServices{
		trades: trd,
		xtb:    xtb,
	}
}
