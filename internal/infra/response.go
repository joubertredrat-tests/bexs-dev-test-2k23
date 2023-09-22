package infra

type RequestValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

type CreatePartnerResponse struct {
	ID          string `json:"id"`
	TradingName string `json:"trading_name"`
	Document    string `json:"document"`
	Currency    string `json:"currency"`
}
