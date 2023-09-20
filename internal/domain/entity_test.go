package domain_test

import (
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrency(t *testing.T) {
	tests := []struct {
		name             string
		currencyInput    string
		currencyExpected domain.Currency
		errExpected      error
	}{
		{
			name:          "Test with valid currency",
			currencyInput: "USD",
			currencyExpected: domain.Currency{
				Value: "USD",
			},
			errExpected: nil,
		},
		{
			name:             "Test with invalid currency",
			currencyInput:    "BRL",
			currencyExpected: domain.Currency{},
			errExpected:      domain.NewErrInvalidCurrency("BRL", domain.Currencies()),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			currencyGot, errGot := domain.NewCurrency(test.currencyInput)

			assert.Equal(t, test.currencyExpected, currencyGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}
