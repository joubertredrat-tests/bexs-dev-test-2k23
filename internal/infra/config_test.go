package infra_test

import (
	"fmt"
	"joubertredrat/bexs-dev-test-2k23/internal/infra"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	apiHostExpected := "0.0.0.0"
	apiPortExpected := "9007"
	databaseHostExpected := "127.0.0.1"
	databasePortExpected := "3306"
	databaseNameExpected := "dbname"
	databaseUserExpected := "user"
	databasePasswordExpected := "password"
	paymentDuplicatedSecondsExpected := uint(60)
	rateUsdExpected := float64(4.75)
	RateEurExpected := float64(5.50)
	RateGbpExpected := float64(6.35)

	os.Setenv("API_HOST", apiHostExpected)
	os.Setenv("API_PORT", apiPortExpected)
	os.Setenv("DATABASE_HOST", databaseHostExpected)
	os.Setenv("DATABASE_PORT", databasePortExpected)
	os.Setenv("DATABASE_NAME", databaseNameExpected)
	os.Setenv("DATABASE_USER", databaseUserExpected)
	os.Setenv("DATABASE_PASSWORD", databasePasswordExpected)
	os.Setenv("PAYMENT_DUPLICATED_SECONDS", fmt.Sprintf("%d", paymentDuplicatedSecondsExpected))
	os.Setenv("RATE_USD", fmt.Sprintf("%v", rateUsdExpected))
	os.Setenv("RATE_EUR", fmt.Sprintf("%v", RateEurExpected))
	os.Setenv("RATE_GBP", fmt.Sprintf("%v", RateGbpExpected))

	config, _ := infra.NewConfig()

	assert.Equal(t, apiHostExpected, config.ApiHost)
	assert.Equal(t, apiPortExpected, config.ApiPort)
	assert.Equal(t, databaseHostExpected, config.DatabaseHost)
	assert.Equal(t, databasePortExpected, config.DatabasePort)
	assert.Equal(t, databaseNameExpected, config.DatabaseName)
	assert.Equal(t, databaseUserExpected, config.DatabaseUser)
	assert.Equal(t, databasePasswordExpected, config.DatabasePassword)
	assert.Equal(t, paymentDuplicatedSecondsExpected, config.PaymentDuplicatedSeconds)
	assert.Equal(t, rateUsdExpected, config.RateUsd)
	assert.Equal(t, RateEurExpected, config.RateEur)
	assert.Equal(t, RateGbpExpected, config.RateGbp)
}
