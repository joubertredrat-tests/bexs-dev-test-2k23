package infra

import (
	"context"
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
)

type ExchangeStatic struct {
}

func NewExchangeStatic() ExchangeStatic {
	return ExchangeStatic{}
}

func (e ExchangeStatic) Convert(ctx context.Context, amount domain.Amount, currency domain.Currency) (domain.Amount, error) {
	return domain.Amount{}, nil
}
