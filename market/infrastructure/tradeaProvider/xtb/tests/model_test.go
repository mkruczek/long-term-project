package tests

import (
	"market/market/domain"
	"testing"
	"time"
)

func Test_ConvertXtbCsvToDomainModel(t *testing.T) {

	testCases := []struct {
		name     string
		xtbCsv   string
		expected domain.Trade
	}{
		{"EURUSD", "12345678,EURUSD,Buy,2020-01-01 00:00:00,1.12345,2020-01-01 00:00:00,1.23456,1000,1000", domain.Trade{
			ID:         "12345678",
			Symbol:     "EURUSD",
			OpenPrice:  domain.Price{Value: 1.12345, Coefficient: 10000},
			OpenTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			ClosePrice: domain.Price{Value: 1.23456, Coefficient: 10000},
			CloseTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Profit:     1111,
			ExternalID: "12345678",
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := convertXtbCsvToDomainModel(tc.xtbCsv)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if actual != tc.expected {
				t.Errorf("expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}

}
