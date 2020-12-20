package domain

import (
	"database/sql"

	/**/
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/siwonpawel/cash-gopher-outside/banking/errs"
	"github.com/siwonpawel/cash-gopher-outside/banking/logger"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error

	var customers []Customer
	if status == "" {
		findAllSQL := "select * from customers"
		err = d.client.Select(&customers, findAllSQL)
	} else {
		findAllWithStatus := "select * from customers where status = ?"
		err = d.client.Select(&customers, findAllWithStatus, status)
	}

	if err != nil {
		logger.Error("Error quering customer table: " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("No data found")
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	sqlQuery := "select * from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, sqlQuery, id)
	if err != nil {
		logger.Error("Error while scanning customer: " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}

		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{client: dbClient}
}
