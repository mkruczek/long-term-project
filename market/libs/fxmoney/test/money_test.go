package test

import (
	"market/market/libs/fxmoney"
	"testing"
)

func Test_CreateNewPrice_Float64(t *testing.T) {

	testCases := []struct {
		name             string
		amount           float64
		currency         string
		expectedAmount   int
		expectedCurrency string
	}{
		{name: "USD", amount: 1.28666, currency: "USD", expectedAmount: 128666, expectedCurrency: "USD"},
		{name: "USD_with_0", amount: 1.0, currency: "USD", expectedAmount: 100000, expectedCurrency: "USD"},
		{name: "JPY", amount: 94.250, currency: "JPY", expectedAmount: 94250, expectedCurrency: "JPY"},
		{name: "JPY_123.123", amount: 123.123, currency: "JPY", expectedAmount: 123123, expectedCurrency: "JPY"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := fxmoney.NewPrice(tc.amount, tc.currency)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if actual.Amount != tc.expectedAmount {
				t.Errorf("expected %v, got %v", tc.expectedAmount, actual.Amount)
			}
			if actual.Currency != tc.expectedCurrency {
				t.Errorf("expected %v, got %v", tc.expectedCurrency, actual.Currency)
			}
		})
	}
}

func Test_CreateNewPrice_String(t *testing.T) {

	testCases := []struct {
		name             string
		amount           string
		currency         string
		expectedAmount   int
		expectedCurrency string
	}{
		{name: "USD", amount: "1.28666", currency: "USD", expectedAmount: 128666, expectedCurrency: "USD"},
		{name: "USD_with_0", amount: "1.0", currency: "USD", expectedAmount: 100000, expectedCurrency: "USD"},
		{name: "JPY", amount: "94.250", currency: "JPY", expectedAmount: 94250, expectedCurrency: "JPY"},
		{name: "JPY_123.123", amount: "123.123", currency: "JPY", expectedAmount: 123123, expectedCurrency: "JPY"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := fxmoney.NewPrice(tc.amount, tc.currency)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if actual.Amount != tc.expectedAmount {
				t.Errorf("expected %v, got %v", tc.expectedAmount, actual.Amount)
			}
			if actual.Currency != tc.expectedCurrency {
				t.Errorf("expected %v, got %v", tc.expectedCurrency, actual.Currency)
			}
		})
	}
}
