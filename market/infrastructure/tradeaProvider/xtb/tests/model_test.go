package tests

import (
	"github.com/google/go-cmp/cmp"
	"market/market/domain"
	"market/market/infrastructure/tradeaProvider/xtb"
	"testing"
	"time"
)

func Test_ConvertXtbCsvToDomainModel(t *testing.T) {

	testCases := []struct {
		name     string
		xtbCsv   xtb.CSV
		expected domain.Trade
	}{
		{name: "EURUSD-BUY",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Buy",
				OpenTime:   "2020-01-01 00:00:00",
				OpenPrice:  1.12345,
				CloseTime:  "2020-01-01 00:00:00",
				ClosePrice: 1.23456,
				Profit:     1000,
				NetProfit:  1000,
			},
			expected: domain.Trade{
				ID:         "12345678",
				Symbol:     "EURUSD",
				OpenPrice:  domain.Price{Value: 1.12345, Coefficient: 1000},
				OpenTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice: domain.Price{Value: 1.23456, Coefficient: 1000},
				CloseTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:     111,
				ExternalID: "12345678",
			}},
		{name: "USDJPY-BUY",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "USDJPY",
				Type:       "Buy",
				OpenTime:   "2020-01-01 00:00:00",
				OpenPrice:  123.345,
				CloseTime:  "2020-01-01 00:00:00",
				ClosePrice: 123.456,
				Profit:     1000,
				NetProfit:  1000,
			},
			expected: domain.Trade{
				ID:         "12345678",
				Symbol:     "USDJPY",
				OpenPrice:  domain.Price{Value: 112.345, Coefficient: 100},
				OpenTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice: domain.Price{Value: 1123.456, Coefficient: 100},
				CloseTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:     111,
			}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.xtbCsv.ToDomainModel()
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if actual != tc.expected {
				t.Errorf("expected is 	different from actual: %v,", cmp.Diff(tc.expected, actual))
			}
		})
	}
}
