package tests

import (
	"context"
	"market/market/domain"
	"market/market/infrastructure/log"
	"testing"
	"time"
)

type dummyGetter struct {
	list []domain.Trade
}

func (d dummyGetter) List(ctx context.Context) ([]domain.Trade, error) {
	return d.list, nil
}

func (d dummyGetter) Get(ctx context.Context, id string) (domain.Trade, error) {
	for _, t := range d.list {
		if t.ID == id {
			return t, nil
		}
	}
	return domain.Trade{}, nil
}

func TestProfit(t *testing.T) {

	log.Init("debug")

	list := []domain.Trade{
		{
			Profit:    10,
			OpenTime:  time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			CloseTime: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Profit:    10,
			OpenTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			CloseTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Profit:    10,
			OpenTime:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			CloseTime: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	dummyGetter := dummyGetter{list: list}

	sut := domain.NewService(dummyGetter)

	got, err := sut.Profit(context.Background(),
		time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 12, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if want := 10; want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}
