package fxmoney

import (
	"fmt"
	"strconv"
	"strings"
)

type Price struct {
	// Amount is the amount of money in the POINTS (not pips) - smallest unit of the currency at fx market.
	// ex. USD -> 1/100000 ; JPY -> 1/1000
	Amount int

	// Currency is the currency of the price.
	Currency string

	// coefficient is the coefficient for the currency. will be set automatically based on the currency.
	coefficient int
}

// Convertible is a generic type for all ints and floats.
type Convertible interface {
	~float64 | ~float32 | ~string
}

// NewPrice creates a new price.
func NewPrice[C Convertible](amount C, currency string) (Price, error) {

	coefficient := getCoefficient(currency)

	stringAmount := fmt.Sprintf("%v", amount)

	floatAmount, err := strconv.ParseFloat(stringAmount, 64)
	if err != nil {
		return Price{}, err
	}

	decimalsNumbers := strings.Count(strconv.Itoa(coefficient), "0")
	stringAmount = fmt.Sprintf("%."+strconv.Itoa(decimalsNumbers)+"f", floatAmount)

	stringAmount = strings.ReplaceAll(stringAmount, ".", "")

	amountInt, err := strconv.Atoi(stringAmount)
	if err != nil {
		return Price{}, err
	}

	return Price{
		Amount:      amountInt,
		Currency:    currency,
		coefficient: coefficient,
	}, nil
}

func (p Price) Subtract(price Price) Price {
	return Price{
		Amount:      p.Amount - price.Amount,
		Currency:    p.Currency,
		coefficient: p.coefficient,
	}
}

func (p Price) Add(price Price) Price {
	return Price{
		Amount:      p.Amount + price.Amount,
		Currency:    p.Currency,
		coefficient: p.coefficient,
	}
}

func getCoefficient(symbol string) int {
	if strings.Contains(symbol, "JPY") {
		return 1000
	}
	return 100000
}
