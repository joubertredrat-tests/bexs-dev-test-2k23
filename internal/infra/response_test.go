package infra_test

import (
	"encoding/json"
	"fmt"
	"joubertredrat/bexs-dev-test-2k23/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestValidationError(t *testing.T) {
	fieldExpected := "name"
	reasonExpected := "required"

	jsonExpected := []byte(fmt.Sprintf(
		`{"field":"%s","reason":"%s"}`,
		fieldExpected,
		reasonExpected,
	))

	response := infra.RequestValidationError{
		Field:  fieldExpected,
		Reason: reasonExpected,
	}
	jsonGot, _ := json.Marshal(response)

	assert.Equal(t, jsonExpected, jsonGot)
}

func TestPartnerResponse(t *testing.T) {
	idExpected := "1"
	tradingNameExpected := "International Ecommerce"
	documentExpected := "1284498339812/0001"
	currencyExpected := "USD"

	jsonExpected := []byte(fmt.Sprintf(
		`{"id":"%s","trading_name":"%s","document":"%s","currency":"%s"}`,
		idExpected,
		tradingNameExpected,
		documentExpected,
		currencyExpected,
	))

	response := infra.PartnerResponse{
		ID:          idExpected,
		TradingName: tradingNameExpected,
		Document:    documentExpected,
		Currency:    currencyExpected,
	}
	jsonGot, _ := json.Marshal(response)

	assert.Equal(t, jsonExpected, jsonGot)
}

func TestPaymentResponse(t *testing.T) {
	idExpected := "01HAW44PR1XK7B027RSFE8SAAY"
	partnerIDExpected := "1"
	amountExpected := "99.05"
	foreignAmountExpected := "470.49"
	consumerNameExpected := "Oliver Tsubasa"
	consumerNationalIDExpected := "30243434597"

	jsonExpected := []byte(fmt.Sprintf(
		`{"id":"%s","partner_id":"%s","amount":"%s","foreign_amount":"%s","consumer":{"name":"%s","national_id":"%s"}}`,
		idExpected,
		partnerIDExpected,
		amountExpected,
		foreignAmountExpected,
		consumerNameExpected,
		consumerNationalIDExpected,
	))

	response := infra.PaymentResponse{
		ID:            idExpected,
		PartnerID:     partnerIDExpected,
		Amount:        amountExpected,
		ForeignAmount: foreignAmountExpected,
		Consumer: infra.ConsumerResponse{
			Name:       consumerNameExpected,
			NationalID: consumerNationalIDExpected,
		},
	}
	jsonGot, _ := json.Marshal(response)

	assert.Equal(t, jsonExpected, jsonGot)
}
