package tests

import (
	"context"
	"market/market/domain"
	"market/market/domain/tradeProvider/xtb"
	"testing"
)

type dummyRepository map[string]domain.Trade

func (d dummyRepository) Insert(ctx context.Context, trade domain.Trade) error {
	d[trade.ID] = trade
	return nil
}

func (d dummyRepository) Read(ctx context.Context, id string) (domain.Trade, error) {
	return d[id], nil
}

func (d dummyRepository) Update(ctx context.Context, trade domain.Trade) error {
	d[trade.ID] = trade
	return nil
}

func (d dummyRepository) Delete(ctx context.Context, id string) error {
	delete(d, id)
	return nil
}

func Test_Provider(t *testing.T) {

	ctx := context.Background()

	traderIDUnderTest := "1037786004"

	data := []*xtb.CSV{
		{Position: traderIDUnderTest, Symbol: "USDCHF", Type: "Buy Stop", OpenTime: "02.08.2023 15:17:31", OpenPrice: 0.87931, CloseTime: "02.08.2023 16:38:00", ClosePrice: 0.87828, Profit: -4.78, NetProfit: -4.78},
		{Position: "1037771461", Symbol: "AUDJPY", Type: "Sell Stop", OpenTime: "02.08.2023 15:08:07", OpenPrice: 93.986, CloseTime: "02.08.2023 15:50:10", ClosePrice: 94.007, Profit: -0.6, NetProfit: -0.6},
		{Position: "1037771460", Symbol: "USDCAD", Type: "Buy", OpenTime: "02.08.2023 08:55:10", OpenPrice: 1.33017, CloseTime: "02.08.2023 13:06:33", ClosePrice: 1.33031, Profit: 0.43, NetProfit: 0.43},
		{Position: "1037771459", Symbol: "AUDUSD", Type: "Sell", OpenTime: "01.08.2023 08:58:03", OpenPrice: 0.66587, CloseTime: "01.08.2023 10:31:03", ClosePrice: 0.66335, Profit: 10.14, NetProfit: 10.14},
	}

	dummyRepo := dummyRepository{}

	p := xtb.NewProvider(dummyRepo)

	err := p.Insert(ctx, data)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if want, got := 4, len(dummyRepo); want != got {
		t.Errorf("want %d, got %d", want, got)
	}

	read, err := p.Read(ctx, traderIDUnderTest)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if want, got := data[0].Position, read.ID; want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
