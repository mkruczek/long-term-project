package xtb

import (
	"market/market/domain"
	"strings"
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
	//	Profit
	Profit float64 `csv:"Profit"`
	//	Net profit
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

	coefficient := getCoefficient(csv.Symbol)

	openPrice := domain.Price{Value: csv.OpenPrice, Coefficient: coefficient}
	closePrice := domain.Price{Value: csv.ClosePrice, Coefficient: coefficient}

	return domain.Trade{
		ID:         csv.Position,
		Symbol:     csv.Symbol,
		OpenPrice:  openPrice,
		OpenTime:   openTime,
		ClosePrice: closePrice,
		CloseTime:  closeTime,
		Profit:     calculateProfit(openPrice, closePrice),
		ExternalID: csv.Position,
	}, nil
}

func calculateProfit(openPrice, closePrice domain.Price) int {
	return int((closePrice.Value - openPrice.Value) * float64(openPrice.Coefficient))
}

func getCoefficient(symbol string) int {
	if strings.Contains(symbol, "JPY") {
		return 100
	}
	return 1000
}

func parseTime(t string) (time.Time, error) {
	result, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return time.Time{}, err
	}
	return result, nil
}
