package server

import (
	"market/market/domain/tradeProvider/xtb"
)

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
