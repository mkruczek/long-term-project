package xtb

import (
	"fmt"
	"market/market/domain"
	"market/market/libs/fxmoney"
	"time"
)

type CSV struct {
	//	Position - ID of the position
	Position string `csv:"Position"`
	//	Symbol
	// e.g. EURUSD, USDJPY etc.
	Symbol string `csv:"Symbol"`
	//	Type - Type of the position
	// e.g. Buy, Sell
	Type string `csv:"Type"`
	//	Open time
	// format: 2006-01-02 15:04:05
	OpenTime string `csv:"Open time"`
	//	Open price
	OpenPrice float64 `csv:"Open price"`
	//	Close time
	// format: 2006-01-02 15:04:05
	CloseTime string `csv:"Close time"`
	//	Close price
	ClosePrice float64 `csv:"Close price"`
	//	Profit in currency
	Profit float64 `csv:"Profit"`
	//	Net profit in currency
	NetProfit float64 `csv:"Net profit"`
}

func (csv CSV) ToDomainModel() (domain.Trade, error) {

	openTime, err := parseTime(csv.OpenTime)
	if err != nil {
		return domain.Trade{}, err
	}

	closeTime, err := parseTime(csv.CloseTime)
	if err != nil {
		return domain.Trade{}, err
	}

	currency := csv.Symbol[3:]

	openPrice, err := fxmoney.NewPrice(csv.OpenPrice, currency)
	if err != nil {
		return domain.Trade{}, err
	}

	closePrice, err := fxmoney.NewPrice(csv.ClosePrice, currency)
	if err != nil {
		return domain.Trade{}, err
	}

	var tradeSide domain.TradeSide
	switch csv.Type {
	case "Buy", "Buy Stop", "Buy Limit":
		tradeSide = domain.Buy
	case "Sell", "Sell Stop", "Sell Limit":
		tradeSide = domain.Sell
	default:
		return domain.Trade{}, fmt.Errorf("unknown trade type: %s", csv.Type)
	}

	result := domain.Trade{
		ID:         csv.Position,
		Symbol:     csv.Symbol,
		TradeSide:  tradeSide,
		OpenPrice:  openPrice,
		OpenTime:   openTime,
		ClosePrice: closePrice,
		CloseTime:  closeTime,
		ExternalID: csv.Position,
	}

	result.CalculateProfit()
	result.CalculateSimplifiedResult()

	return result, nil
}

func parseTime(t string) (time.Time, error) {
	result, err := time.Parse("02.01.2006 15:04:05", t)
	if err != nil {
		return time.Time{}, err
	}
	return result, nil
}
