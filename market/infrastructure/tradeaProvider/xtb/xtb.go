package xtb

import "market/market/domain"

type Provider struct {
}

func NewProvider() Provider {
	return Provider{}
}

func (p Provider) LoadTrades() ([]domain.GenericModel, error) {
	return nil, nil
}
