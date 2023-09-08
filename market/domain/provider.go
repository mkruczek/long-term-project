package domain

import "context"

type ToModel interface {
	ToDomainModel() (Trade, error)
}

type inserter interface {
	//todo -> change to bulk insert
	Insert(ctx context.Context, trade Trade) error
}

type Provider[K ToModel] struct {
	repo inserter
}

func NewProvider[K ToModel](r inserter) Provider[K] {
	return Provider[K]{repo: r}
}

func (p Provider[K]) Insert(ctx context.Context, data []K) error {
	for _, v := range data {
		dm, err := v.ToDomainModel()
		if err != nil {
			return err
		}
		err = p.repo.Insert(ctx, dm)
		if err != nil {
			//todo transaction or bulk
			return err
		}
	}
	return nil
}
