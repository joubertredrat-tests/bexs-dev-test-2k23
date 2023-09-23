package infra_test

import (
	"encoding/json"
	"fmt"
	"joubertredrat/bexs-dev-test-2k23/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePartnerRequest(t *testing.T) {
	idExpected := "1"
	tradingNameExpected := "International Ecommerce"
	documentExpected := "1284498339812/0001"
	currencyExpected := "USD"

	jsonData := []byte(fmt.Sprintf(
		`{"id":"%s","trading_name":"%s","document":"%s","currency":"%s"}`,
		idExpected,
		tradingNameExpected,
		documentExpected,
		currencyExpected,
	))

	var request infra.CreatePartnerRequest
	json.Unmarshal(jsonData, &request)

	assert.Equal(t, idExpected, request.ID)
	assert.Equal(t, tradingNameExpected, request.TradingName)
	assert.Equal(t, documentExpected, request.Document)
	assert.Equal(t, currencyExpected, request.Currency)
}

func TestCreatePaymentRequest(t *testing.T) {
	partnerIDExpected := "1"
	amountExpected := "99.05"
	consumerNameExpected := "Oliver Tsubasa"
	consumerNationalIDExpected := "30243434597"

	jsonData := []byte(fmt.Sprintf(
		`{"partner_id":"%s","amount":"%s","consumer":{"name":"%s","national_id":"%s"}}`,
		partnerIDExpected,
		amountExpected,
		consumerNameExpected,
		consumerNationalIDExpected,
	))

	var request infra.CreatePaymentRequest
	json.Unmarshal(jsonData, &request)

	assert.Equal(t, partnerIDExpected, request.PartnerID)
	assert.Equal(t, amountExpected, request.Amount)
	assert.Equal(t, consumerNameExpected, request.Consumer.Name)
	assert.Equal(t, consumerNationalIDExpected, request.Consumer.NationalID)
}
