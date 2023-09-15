package domain

import "context"

type toModel interface {
	ToDomainModel() (Trade, error)
}

type inserter interface {
	BulkInsert(ctx context.Context, trades []Trade) error
}

type Provider[K toModel] struct {
	repo inserter
}

func NewProvider[K toModel](r inserter) Provider[K] {
	return Provider[K]{repo: r}
}

func (p Provider[K]) BulkInsert(ctx context.Context, data []K) error {

	trades := make([]Trade, len(data))
	for i, v := range data {
		dm, err := v.ToDomainModel()
		if err != nil {
			return err
		}
		trades[i] = dm
	}
	err := p.repo.BulkInsert(ctx, trades)
	if err != nil {
		return err
	}
	return nil
}
