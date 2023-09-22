package application

type UsecaseCreatePartnerInput struct {
	ID          string
	TradingName string
	Document    string
	Currency    string
}

type UsecaseCreatePaymentInput struct {
	PartnerID          string
	Amount             string
	ConsumerName       string
	ConsumerNationalID string
}
