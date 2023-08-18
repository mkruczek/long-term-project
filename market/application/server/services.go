package server

import "market/market/infrastructure/tradeaProvider/xtb"

// appServices - services for application
type appServices struct {
	XtbProvider xtb.Provider
}

// NewAppServices - create new appServices
func NewAppServices(xtbProvider xtb.Provider) appServices {
	return appServices{
		XtbProvider: xtbProvider,
	}
}
