package domain

import (
	"time"

	"github.com/siwonpawel/cash-gopher-outside/banking/dto"
	"github.com/siwonpawel/cash-gopher-outside/banking/errs"
)

type Account struct {
	AccountID   string  `db:"account_id"`
	CustomerID  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{
		AccountId: a.AccountID,
	}
}

func FromNewAccountRequest(a dto.NewAccountRequest) (*Account, *errs.AppError) {

	return &Account{
		AccountID:   "",
		CustomerID:  a.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      "1",
	}, nil
}

type AccountRepository interface {
	Save(*Account) (*Account, *errs.AppError)
}
