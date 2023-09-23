package infra

type RequestValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

type PartnerResponse struct {
	ID          string `json:"id"`
	TradingName string `json:"trading_name"`
	Document    string `json:"document"`
	Currency    string `json:"currency"`
}

type PaymentResponse struct {
	ID            string           `json:"id"`
	PartnerID     string           `json:"partner_id"`
	Amount        string           `json:"amount"`
	ForeignAmount string           `json:"foreign_amount"`
	Consumer      ConsumerResponse `json:"consumer"`
}

type ConsumerResponse struct {
	Name       string `json:"name"`
	NationalID string `json:"national_id"`
}
