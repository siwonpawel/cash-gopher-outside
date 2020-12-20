package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/siwonpawel/cash-gopher-outside/banking/errs"
	"github.com/siwonpawel/cash-gopher-outside/banking/logger"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return TransactionRepositoryDB{client: db}
}

func (repo TransactionRepositoryDB) AddTransaction(t Transaction) (*Account, *Transaction, *errs.AppError) {

	var acc Account
	selectAccountInformation := "SELECT * FROM accounts WHERE account_id = ?"
	if err := repo.client.Get(&acc, selectAccountInformation, t.AccountID); err != nil {
		logger.Error("Error quering account " + err.Error())
		if err == sql.ErrNoRows {
			return nil, nil, errs.NewNotFoundError("Account not found")
		}

		return nil, nil, errs.NewUnexpectedError("Database querying problem")
	}

	if t.IsWithdrawal() && t.Amount > acc.Amount {
		return nil, nil, errs.NewValidationError("You cannot w")
	}

	var updateAccountBalance string
	if t.IsWithdrawal() {
		updateAccountBalance = "UPDATE accounts SET amount = amount - ? WHERE account_id = ?"
		acc.Amount -= t.Amount
	} else {
		updateAccountBalance = "UPDATE accounts SET amount = amount + ? WHERE account_id = ?"
		acc.Amount += t.Amount
	}

	tx, err := repo.client.DB.Begin()
	if err != nil {
		logger.Error("Error when starting transaction" + err.Error())
		return nil, nil, errs.NewUnexpectedError("Database querying problem")
	}

	_, err = tx.Exec(updateAccountBalance, t.Amount, t.AccountID)
	if err != nil {
		logger.Error("Exception updating account: " + err.Error())
		tx.Rollback()
		return nil, nil, errs.NewUnexpectedError("Error updating account with transaction " + err.Error())
	}

	insertTransacion := "INSERT INTO transactions(account_id, amount, transaction_type, transaction_date) VALUES(?, ?, ?, ?)"
	result, err := tx.Exec(insertTransacion, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Exception inserting transaction: " + err.Error())
		tx.Rollback()
		return nil, nil, errs.NewUnexpectedError("Error creating transaction " + err.Error())
	}

	tID, err := result.LastInsertId()
	t.TransactionID = tID

	err = tx.Commit()
	if err != nil {
		logger.Error("Error commiting changes: " + err.Error())
		tx.Rollback()
		return nil, nil, errs.NewUnexpectedError("Error with update")
	}

	return &acc, &t, nil
}
