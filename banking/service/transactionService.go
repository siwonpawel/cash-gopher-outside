package service

import (
	"strconv"
	"time"

	"github.com/siwonpawel/cash-gopher-outside/banking/domain"
	"github.com/siwonpawel/cash-gopher-outside/banking/dto"
	"github.com/siwonpawel/cash-gopher-outside/banking/errs"
)

type DefaultTransactionService struct {
	transactionRepo domain.TransactionRepository
	accountRepo     domain.AccountRepository
}

func NewTransactionService(transactionRepository domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{transactionRepo: transactionRepository}
}

func (service DefaultTransactionService) AddTransaction(tr dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {

	if err := tr.Validate(); err != nil {
		return nil, err
	}

	//cID, _ := strconv.Atoi(tr.CustomerID)
	aID, _ := strconv.Atoi(tr.AccountID)

	transaction := domain.Transaction{
		AccountID:       aID,
		TransactionType: tr.TransactionType,
		Amount:          tr.Amount,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	acc, trn, err := service.transactionRepo.AddTransaction(transaction)
	if err != nil {
		return nil, err
	}

	return &dto.TransactionResponse{
		TransactionID: trn.TransactionID,
		Balance:       acc.Amount,
	}, nil
}
