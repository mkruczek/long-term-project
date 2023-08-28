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
		{name: "EURUSD-BUY-1pointProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Buy",
				OpenTime:   "2020-01-01 00:00:00",
				OpenPrice:  1.00005,
				CloseTime:  "2020-01-01 00:00:00",
				ClosePrice: 1.00006,
				Profit:     0, //for single point (1/10 pip) profit is 0
				NetProfit:  0,
			},
			expected: domain.Trade{
				ID:         "12345678",
				Symbol:     "EURUSD",
				TradeSide:  domain.Buy,
				OpenPrice:  domain.Price{Value: 1.00005, Coefficient: 100000},
				OpenTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice: domain.Price{Value: 1.00006, Coefficient: 100000},
				CloseTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:     1,
				ExternalID: "12345678",
			}},
		{name: "USDJPY-BUY-1pointProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "USDJPY",
				Type:       "Buy",
				OpenTime:   "2020-01-01 00:00:00",
				OpenPrice:  123.005,
				CloseTime:  "2020-01-01 00:00:00",
				ClosePrice: 123.006,
				Profit:     1,
				NetProfit:  1,
			},
			expected: domain.Trade{
				ID:         "12345678",
				Symbol:     "USDJPY",
				TradeSide:  domain.Buy,
				OpenPrice:  domain.Price{Value: 123.005, Coefficient: 1000},
				OpenTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice: domain.Price{Value: 123.006, Coefficient: 1000},
				CloseTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:     1,
				ExternalID: "12345678",
			}},
		{name: "EURUSD-BUY-1pointLoss",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Buy",
				OpenTime:   "2020-01-01 00:00:00",
				OpenPrice:  1.00005,
				CloseTime:  "2020-01-01 00:00:00",
				ClosePrice: 1.00004,
				Profit:     0, //for single point (1/10 pip) loss is 0
				NetProfit:  0,
			},
			expected: domain.Trade{
				ID:         "12345678",
				Symbol:     "EURUSD",
				TradeSide:  domain.Buy,
				OpenPrice:  domain.Price{Value: 1.00005, Coefficient: 100000},
				OpenTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice: domain.Price{Value: 1.00004, Coefficient: 100000},
				CloseTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:     -1,
				ExternalID: "12345678",
			},
		},
		{name: "USDJPY-BUY-1pointLoss",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "USDJPY",
				Type:       "Buy",
				OpenTime:   "2020-01-01 00:00:00",
				OpenPrice:  123.005,
				CloseTime:  "2020-01-01 00:00:00",
				ClosePrice: 123.004,
				Profit:     -1,
				NetProfit:  -1,
			},
			expected: domain.Trade{
				ID:         "12345678",
				Symbol:     "USDJPY",
				TradeSide:  domain.Buy,
				OpenPrice:  domain.Price{Value: 123.005, Coefficient: 1000},
				OpenTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice: domain.Price{Value: 123.004, Coefficient: 1000},
				CloseTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:     -1,
				ExternalID: "12345678",
			},
		},
		{name: "EURUSD-Buy-1PipProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Buy",
				OpenTime:   "2020-01-01 00:00:00",
				OpenPrice:  1.00005,
				CloseTime:  "2020-01-01 00:00:00",
				ClosePrice: 1.00015,
				Profit:     10, //for single pip profit
				NetProfit:  10,
			},
			expected: domain.Trade{
				ID:         "12345678",
				Symbol:     "EURUSD",
				TradeSide:  domain.Buy,
				OpenPrice:  domain.Price{Value: 1.00005, Coefficient: 100000},
				OpenTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice: domain.Price{Value: 1.00015, Coefficient: 100000},
				CloseTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:     10,
				ExternalID: "12345678",
			},
		},
		{name: "EURUSD-Buy-1PipLos",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Buy",
				OpenTime:   "2020-01-01 00:00:00",
				OpenPrice:  1.00015,
				CloseTime:  "2020-01-01 00:00:00",
				ClosePrice: 1.00005,
				Profit:     -10, //for single pip profit
				NetProfit:  -10,
			},
			expected: domain.Trade{
				ID:         "12345678",
				Symbol:     "EURUSD",
				TradeSide:  domain.Buy,
				OpenPrice:  domain.Price{Value: 1.00015, Coefficient: 100000},
				OpenTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice: domain.Price{Value: 1.00005, Coefficient: 100000},
				CloseTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:     -10,
				ExternalID: "12345678",
			},
		},
		{name: "EURJPY-Buy-1PipProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURJPY",
				Type:       "Buy",
				OpenTime:   "2020-01-01 00:00:00",
				OpenPrice:  123.005,
				CloseTime:  "2020-01-01 00:00:00",
				ClosePrice: 123.015,
				Profit:     10, //for single pip profit
				NetProfit:  10,
			},
			expected: domain.Trade{
				ID:         "12345678",
				Symbol:     "EURJPY",
				TradeSide:  domain.Buy,
				OpenPrice:  domain.Price{Value: 123.005, Coefficient: 1000},
				OpenTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice: domain.Price{Value: 123.015, Coefficient: 1000},
				CloseTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:     10,
				ExternalID: "12345678",
			},
		},
		{name: "EURJPY-Buy-1PipLos",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURJPY",
				Type:       "Buy",
				OpenTime:   "2020-01-01 00:00:00",
				OpenPrice:  123.015,
				CloseTime:  "2020-01-01 00:00:00",
				ClosePrice: 123.005,
				Profit:     -10, //for single pip profit
				NetProfit:  -10,
			},
			expected: domain.Trade{
				ID:         "12345678",
				Symbol:     "EURJPY",
				TradeSide:  domain.Buy,
				OpenPrice:  domain.Price{Value: 123.015, Coefficient: 1000},
				OpenTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice: domain.Price{Value: 123.005, Coefficient: 1000},
				CloseTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:     -10,
				ExternalID: "12345678",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
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
