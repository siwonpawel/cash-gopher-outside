package dto

type TransactionResponse struct {
	TransactionID int64   `json:"transaction_id"`
	Balance       float64 `json:"balance"`
}
