package server

import (
	"market/market/application/statistic/useCase"
)

// appServices - services for application
type appServices struct {
	sts useCase.StatisticGetter
}

// NewAppServices - create new appServices
func NewAppServices(sts useCase.StatisticGetter) appServices {
	return appServices{
		sts: sts,
	}
}
