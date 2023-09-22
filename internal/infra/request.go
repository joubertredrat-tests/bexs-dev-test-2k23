package infra

type CreatePartnerRequest struct {
	ID          string `json:"id" binding:"required"`
	TradingName string `json:"trading_name" binding:"required"`
	Document    string `json:"document" binding:"required"`
	Currency    string `json:"currency" binding:"required"`
}
