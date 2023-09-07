package xtb

import (
	"context"
	"market/market/domain"
)

type repository interface {
	Insert(ctx context.Context, trade domain.Trade) error
	Read(ctx context.Context, id string) (domain.Trade, error)
	Update(ctx context.Context, trade domain.Trade) error
	Delete(ctx context.Context, id string) error
}

type Provider struct {
	repo repository
}

func NewProvider(r repository) Provider {
	return Provider{repo: r}
}

func (p Provider) Insert(ctx context.Context, data []*CSV) error {
	for _, v := range data {
		dm, err := v.ToDomainModel()
		if err != nil {
			return err
		}
		err = p.repo.Insert(ctx, dm)
		if err != nil {
			//todo transaction
			return err
		}
	}
	return nil
}

func (p Provider) Read(ctx context.Context, id string) (domain.Trade, error) {
	return p.repo.Read(ctx, id)
}

func (p Provider) Update(ctx context.Context, trade domain.Trade) error {
	return p.repo.Update(ctx, trade)
}

func (p Provider) Delete(ctx context.Context, id string) error {
	return p.repo.Delete(ctx, id)
}
