package xtb_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"market/market/domain/dataProviders/xtb"
	"market/market/domain/trade"
	"market/market/libs/fxmoney"
	"testing"
	"time"
)

var cmpOpts = cmp.Options{cmpopts.IgnoreFields(fxmoney.Price{}, "coefficient")}

func Test_ConvertXtbCsvToDomainModel_BasicValues(t *testing.T) {

	testCases := []struct {
		name     string
		xtbCsv   xtb.CSV
		expected trade.Trade
	}{
		{name: "EURUSD-Buy-1PointProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Buy",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  1.00005,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 1.00006,
				Profit:     0, //for single point (1/10 pip) profit is 0
				NetProfit:  0,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURUSD",
				TradeSide:        trade.Buy,
				OpenPrice:        fxmoney.Price{Amount: 100005, Currency: "USD"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 100006, Currency: "USD"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           1,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			}},
		{name: "USDJPY-Buy-1PointProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "USDJPY",
				Type:       "Buy",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  123.005,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 123.006,
				Profit:     1,
				NetProfit:  1,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "USDJPY",
				TradeSide:        trade.Buy,
				OpenPrice:        fxmoney.Price{Amount: 123005, Currency: "JPY"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 123006, Currency: "JPY"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           1,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			}},
		{name: "EURUSD-Buy-1PointLoss",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Buy",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  1.00005,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 1.00004,
				Profit:     0, //for single point (1/10 pip) loss is 0
				NetProfit:  0,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURUSD",
				TradeSide:        trade.Buy,
				OpenPrice:        fxmoney.Price{Amount: 100005, Currency: "USD"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 100004, Currency: "USD"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           -1,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "USDJPY-Buy-1PointLoss",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "USDJPY",
				Type:       "Buy",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  123.005,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 123.004,
				Profit:     -1,
				NetProfit:  -1,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "USDJPY",
				TradeSide:        trade.Buy,
				OpenPrice:        fxmoney.Price{Amount: 123005, Currency: "JPY"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 123004, Currency: "JPY"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           -1,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "EURUSD-Buy-1PipProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Buy",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  1.00005,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 1.00015,
				Profit:     10, //for single pip profit
				NetProfit:  10,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURUSD",
				TradeSide:        trade.Buy,
				OpenPrice:        fxmoney.Price{Amount: 100005, Currency: "USD"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 100015, Currency: "USD"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           10,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "EURUSD-Buy-1PipLoss",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Buy",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  1.00015,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 1.00005,
				Profit:     -10, //for single pip profit
				NetProfit:  -10,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURUSD",
				TradeSide:        trade.Buy,
				OpenPrice:        fxmoney.Price{Amount: 100015, Currency: "USD"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 100005, Currency: "USD"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           -10,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "EURJPY-Buy-1PipProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURJPY",
				Type:       "Buy",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  123.005,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 123.015,
				Profit:     10, //for single pip profit
				NetProfit:  10,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURJPY",
				TradeSide:        trade.Buy,
				OpenPrice:        fxmoney.Price{Amount: 123005, Currency: "JPY"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 123015, Currency: "JPY"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           10,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "EURJPY-Buy-1PipLoss",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURJPY",
				Type:       "Buy",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  123.015,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 123.005,
				Profit:     -10, //for single pip profit
				NetProfit:  -10,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURJPY",
				TradeSide:        trade.Buy,
				OpenPrice:        fxmoney.Price{Amount: 123015, Currency: "JPY"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 123005, Currency: "JPY"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           -10,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "EURUSD-Sell-1PointProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Sell",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  1.00015,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 1.00014,
				Profit:     0, //for single point (1/10 pip) profit is 0
				NetProfit:  0,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURUSD",
				TradeSide:        trade.Sell,
				OpenPrice:        fxmoney.Price{Amount: 100015, Currency: "USD"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 100014, Currency: "USD"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           1,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "USDJPY-Sell-1PointProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "USDJPY",
				Type:       "Sell",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  123.015,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 123.014,
				Profit:     1,
				NetProfit:  1,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "USDJPY",
				TradeSide:        trade.Sell,
				OpenPrice:        fxmoney.Price{Amount: 123015, Currency: "JPY"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 123014, Currency: "JPY"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           1,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "EURUSD-Sell-1PointLoss",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Sell",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  1.00015,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 1.00016,
				Profit:     0, //for single point (1/10 pip) loss is 0
				NetProfit:  0,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURUSD",
				TradeSide:        trade.Sell,
				OpenPrice:        fxmoney.Price{Amount: 100015, Currency: "USD"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 100016, Currency: "USD"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           -1,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "USDJPY-Sell-1PointLoss",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "USDJPY",
				Type:       "Sell",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  123.015,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 123.016,
				Profit:     -1,
				NetProfit:  -1,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "USDJPY",
				TradeSide:        trade.Sell,
				OpenPrice:        fxmoney.Price{Amount: 123015, Currency: "JPY"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 123016, Currency: "JPY"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           -1,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			}},
		{name: "EURUSD-Sell-1PipProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Sell",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  1.00015,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 1.00005,
				Profit:     10, //for single pip profit
				NetProfit:  10,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURUSD",
				TradeSide:        trade.Sell,
				OpenPrice:        fxmoney.Price{Amount: 100015, Currency: "USD"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 100005, Currency: "USD"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           10,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "EURUSD-Sell-1PipLoss",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURUSD",
				Type:       "Sell",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  1.00005,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 1.00015,
				Profit:     -10, //for single pip profit
				NetProfit:  -10,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURUSD",
				TradeSide:        trade.Sell,
				OpenPrice:        fxmoney.Price{Amount: 100005, Currency: "USD"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 100015, Currency: "USD"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           -10,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "EURJPY-Sell-1PipProfit",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURJPY",
				Type:       "Sell",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  123.015,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 123.005,
				Profit:     10, //for single pip profit
				NetProfit:  10,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURJPY",
				TradeSide:        trade.Sell,
				OpenPrice:        fxmoney.Price{Amount: 123015, Currency: "JPY"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 123005, Currency: "JPY"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           10,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
			},
		},
		{name: "EURJPY-Sell-1PipLoss",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "EURJPY",
				Type:       "Sell",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  123.005,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 123.015,
				Profit:     -10, //for single pip profit
				NetProfit:  -10,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "EURJPY",
				TradeSide:        trade.Sell,
				OpenPrice:        fxmoney.Price{Amount: 123005, Currency: "JPY"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 123015, Currency: "JPY"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           -10,
				SimplifiedResult: trade.BreakEven,
				ExternalID:       "12345678",
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
			if !cmp.Equal(actual, tc.expected, cmpOpts) {
				t.Errorf("expected is 	different from actual: %v,", cmp.Diff(tc.expected, actual, cmpOpts))
			}
		})
	}
}

func Test_ConvertXtbCsvToDomainModel_RealValues(t *testing.T) {

	testCases := []struct {
		name     string
		xtbCsv   xtb.CSV
		expected trade.Trade
	}{
		{name: "GBPUSD-Buy-Loss",
			xtbCsv: xtb.CSV{
				Position:   "12345678",
				Symbol:     "GBPUSD",
				Type:       "Buy",
				OpenTime:   "01.01.2020 00:00:00",
				OpenPrice:  1.28666,
				CloseTime:  "01.01.2020 00:00:00",
				ClosePrice: 1.28567,
				Profit:     -3.97,
				NetProfit:  -3.97,
			},
			expected: trade.Trade{
				ID:               "12345678",
				Symbol:           "GBPUSD",
				TradeSide:        trade.Buy,
				OpenPrice:        fxmoney.Price{Amount: 128666, Currency: "USD"},
				OpenTime:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ClosePrice:       fxmoney.Price{Amount: 128567, Currency: "USD"},
				CloseTime:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Profit:           -99,
				SimplifiedResult: trade.Loss,
				ExternalID:       "12345678",
			}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tc.xtbCsv.ToDomainModel()
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if !cmp.Equal(actual, tc.expected, cmpOpts) {
				t.Errorf("expected is 	different from actual: %v,", cmp.Diff(tc.expected, actual, cmpOpts))
			}
		})
	}
}
