package domain

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
	tradingName string
	document    string
	currency    Currency
}
