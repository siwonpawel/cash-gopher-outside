package dto

import (
	"github.com/siwonpawel/cash-gopher-outside/banking/errs"
)

type TransactionRequest struct {
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	CustomerID      string  `json:"customer_id"`
	AccountID       string  `json:"account_id"`
}

func (tr TransactionRequest) Validate() *errs.AppError {

	if tr.Amount <= 0 {
		return errs.NewValidationError("Amount cannot be less or equal to 0.00")
	}

	if tr.TransactionType != "withdrawal" && tr.TransactionType != "deposit" {
		return errs.NewValidationError("Transaction type need to be withdrawal or deposit")
	}

	return nil
}
