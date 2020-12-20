package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/siwonpawel/cash-gopher-outside/banking/errs"
	"github.com/siwonpawel/cash-gopher-outside/banking/logger"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a *Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts(customer_id, opening_date, account_type, amount, status) VALUES(?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexptected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Undexpected error from database")
	}

	a.AccountID = strconv.FormatInt(id, 10)
	return a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client: dbClient}
}
