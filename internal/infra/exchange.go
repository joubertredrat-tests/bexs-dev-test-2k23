package infra

import (
	"context"
	"fmt"
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
	"strconv"
)

type ExchangeStatic struct {
	rateUsd float64
	rateEur float64
	rateGbp float64
}

func NewExchangeStatic(ru, re, rg float64) ExchangeStatic {
	return ExchangeStatic{
		rateUsd: ru,
		rateEur: re,
		rateGbp: rg,
	}
}

func (e ExchangeStatic) Convert(ctx context.Context, amount domain.Amount, currency domain.Currency) (domain.Amount, error) {
	rates := e.rates()
	rate, ok := rates[currency.Value]
	if !ok {
		return amount, nil
	}

	v, err := float(amount)
	if err != nil {
		return domain.Amount{}, err
	}

	return domain.NewAmount(fmt.Sprintf("%.2f", v*rate))
}

func (e ExchangeStatic) rates() map[string]float64 {
	return map[string]float64{
		domain.CURRENCY_USD: e.rateUsd,
		domain.CURRENCY_EUR: e.rateEur,
		domain.CURRENCY_GBP: e.rateGbp,
	}
}

func float(a domain.Amount) (float64, error) {
	return strconv.ParseFloat(a.Value, 64)
}
