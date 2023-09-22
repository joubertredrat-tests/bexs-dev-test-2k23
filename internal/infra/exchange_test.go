package infra_test

import (
	"context"
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
	"joubertredrat/bexs-dev-test-2k23/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchangeStatic(t *testing.T) {
	tests := []struct {
		name           string
		rateUsd        float64
		rateEur        float64
		rateGbp        float64
		amount         domain.Amount
		currency       domain.Currency
		amountExpected domain.Amount
		errExpected    error
	}{
		{
			name:    "test convert to USD",
			rateUsd: 4.75,
			rateEur: 5.50,
			rateGbp: 6.35,
			amount: domain.Amount{
				Value: "99.05",
			},
			currency: domain.Currency{
				Value: domain.CURRENCY_USD,
			},
			amountExpected: domain.Amount{
				Value: "470.49",
			},
			errExpected: nil,
		},
		{
			name:    "test convert to EUR",
			rateUsd: 4.75,
			rateEur: 5.50,
			rateGbp: 6.35,
			amount: domain.Amount{
				Value: "99.05",
			},
			currency: domain.Currency{
				Value: domain.CURRENCY_EUR,
			},
			amountExpected: domain.Amount{
				Value: "544.77",
			},
			errExpected: nil,
		},
		{
			name:    "test convert to GBP",
			rateUsd: 4.75,
			rateEur: 5.50,
			rateGbp: 6.35,
			amount: domain.Amount{
				Value: "99.05",
			},
			currency: domain.Currency{
				Value: domain.CURRENCY_GBP,
			},
			amountExpected: domain.Amount{
				Value: "628.97",
			},
			errExpected: nil,
		},
		{
			name:    "test convert to unknown currency",
			rateUsd: 4.75,
			rateEur: 5.50,
			rateGbp: 6.35,
			amount: domain.Amount{
				Value: "99.05",
			},
			currency: domain.Currency{
				Value: "BRL",
			},
			amountExpected: domain.Amount{
				Value: "99.05",
			},
			errExpected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()

			exchange := infra.NewExchangeStatic(test.rateUsd, test.rateEur, test.rateGbp)
			amountGot, errGot := exchange.Convert(ctx, test.amount, test.currency)

			assert.Equal(t, test.amountExpected, amountGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}
