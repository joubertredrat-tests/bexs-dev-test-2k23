package domain_test

import (
	"fmt"
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

func TestPartner(t *testing.T) {
	idExpected := "1"
	tradingNameExpected := "International Ecommerce"
	documentExpected := "1284498339812/0001"
	currencyExpected, _ := domain.NewCurrency(domain.CURRENCY_USD)

	partnerGot := domain.NewPartner(idExpected, tradingNameExpected, documentExpected, currencyExpected)

	assert.Equal(t, idExpected, partnerGot.ID)
	assert.Equal(t, tradingNameExpected, partnerGot.TradingName)
	assert.Equal(t, documentExpected, partnerGot.Document)
	assert.Equal(t, currencyExpected, partnerGot.Currency)
}

func TestAmount(t *testing.T) {
	tests := []struct {
		name           string
		valueInput     string
		amountExpected domain.Amount
		errExpected    error
	}{
		{
			name:       "Test with valid data",
			valueInput: "99.05",
			amountExpected: domain.Amount{
				Value: "99.05",
			},
			errExpected: nil,
		},
		{
			name:           "Test with invalid value without dot",
			valueInput:     "9905",
			amountExpected: domain.Amount{},
			errExpected:    fmt.Errorf("Amount expect valid value, got [ 9905 ]"),
		},
		{
			name:           "Test with invalid value on decimals",
			valueInput:     "990.5",
			amountExpected: domain.Amount{},
			errExpected:    fmt.Errorf("Amount expect valid value, got [ 990.5 ]"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			amountGot, errGot := domain.NewAmount(test.valueInput)

			assert.Equal(t, test.amountExpected, amountGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}

func TestConsumer(t *testing.T) {
	tests := []struct {
		name             string
		nameInput        string
		nationalIDInput  string
		consumerExpected domain.Consumer
		errExpected      error
	}{
		{
			name:            "Test with valid data",
			nameInput:       "Oliver Tsubasa",
			nationalIDInput: "30243434597",
			consumerExpected: domain.Consumer{
				Name:       "Oliver Tsubasa",
				NationalID: "30243434597",
			},
			errExpected: nil,
		},
		{
			name:             "Test with invalid national ID",
			nameInput:        "Oliver Tsubasa",
			nationalIDInput:  "302434597",
			consumerExpected: domain.Consumer{},
			errExpected:      fmt.Errorf("Consumer national ID expect 11 digits, got [ 302434597 ]"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			consumerGot, errGot := domain.NewConsumer(test.nameInput, test.nationalIDInput)

			assert.Equal(t, test.consumerExpected, consumerGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}
