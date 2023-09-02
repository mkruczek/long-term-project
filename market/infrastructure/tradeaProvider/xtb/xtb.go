package xtb

import "market/market/domain"

type Provider struct {
}

func NewProvider() Provider {
	return Provider{}
}

func (p Provider) UpsertTrades(data []*CSV) error {

	models := make([]domain.Trade, len(data))

	var profit int

	for i, v := range data {
		dm, err := v.ToDomainModel()

		profit += dm.Profit

		if err != nil {
			return err
		}
		models[i] = dm
	}

	//todo add to db

	return nil
}
