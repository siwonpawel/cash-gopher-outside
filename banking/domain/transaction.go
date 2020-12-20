package domain

import (
	"github.com/siwonpawel/cash-gopher-outside/banking/errs"
)

const (
	WITHDRAWAL = "withdrawal"
	DEPOSIT    = "deposit"
)

type Transaction struct {
	TransactionID   int64   `db:"transaction_id"`
	AccountID       int     `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transcation_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

type TransactionRepository interface {
	AddTransaction(Transaction) (*Account, *Transaction, *errs.AppError)
}
