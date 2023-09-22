package infra

type CreatePartnerRequest struct {
	ID          string `json:"id" binding:"required"`
	TradingName string `json:"trading_name" binding:"required"`
	Document    string `json:"document" binding:"required"`
	Currency    string `json:"currency" binding:"required"`
}

type CreatePaymentRequest struct {
	PartnerID string          `json:"partner_id" binding:"required"`
	Amount    string          `json:"amount" binding:"required"`
	Consumer  ConsumerRequest `json:"consumer" binding:"required"`
}

type ConsumerRequest struct {
	Name       string `json:"name" binding:"required"`
	NationalID string `json:"national_id" binding:"required"`
}
