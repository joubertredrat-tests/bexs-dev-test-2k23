package domain

import (
	"fmt"
	"regexp"
	"time"
)

const (
	CURRENCY_USD = "USD"
	CURRENCY_EUR = "EUR"
	CURRENCY_GBP = "GBP"
)

type Currency struct {
	Value string
}

func NewCurrency(currency string) (Currency, error) {
	if !contains(currency, Currencies()) {
		return Currency{}, NewErrInvalidCurrency(currency, Currencies())
	}

	return Currency{
		Value: currency,
	}, nil
}

func Currencies() []string {
	return []string{
		CURRENCY_USD,
		CURRENCY_EUR,
		CURRENCY_GBP,
	}
}

type Partner struct {
	ID          string
	TradingName string
	Document    string
	Currency    Currency
}

func NewPartner(ID, tradingName, document string, currency Currency) Partner {
	return Partner{
		ID:          ID,
		TradingName: tradingName,
		Document:    document,
		Currency:    currency,
	}
}

type Amount struct {
	Value string
}

func NewAmount(value string) (Amount, error) {
	if !regexp.MustCompile(`^[\d]{1,}[.][\d]{2}$`).MatchString(value) {
		return Amount{}, fmt.Errorf("Amount expect valid value, got [ %s ]", value)
	}

	return Amount{
		Value: value,
	}, nil
}

type Consumer struct {
	Name       string
	NationalID string
}

func NewConsumer(name, nationalID string) (Consumer, error) {
	if !regexp.MustCompile(`^[\d]{11}$`).MatchString(nationalID) {
		return Consumer{}, fmt.Errorf("Consumer national ID expect 11 digits, got [ %s ]", nationalID)
	}

	return Consumer{
		Name:       name,
		NationalID: nationalID,
	}, nil
}

type Payment struct {
	ID            string
	PartnerID     string
	Amount        Amount
	ForeignAmount Amount
	Consumer      Consumer
	Created       time.Time
}
