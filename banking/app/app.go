package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/siwonpawel/cash-gopher-outside/banking/config"
	"github.com/siwonpawel/cash-gopher-outside/banking/domain"
	"github.com/siwonpawel/cash-gopher-outside/banking/logger"
	"github.com/siwonpawel/cash-gopher-outside/banking/service"
)

func Start() {

	router := mux.NewRouter()

	dbClient := getDbClient()
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient))}
	ah := AccountHandlers{service: service.NewAccountService(domain.NewAccountRepositoryDb(dbClient))}
	th := TransactionHandlers{service: service.NewTransactionService(domain.NewTransactionRepository(dbClient))}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", th.AddTransaction).Methods(http.MethodPost)

	appAdress := config.GetAppAddress()
	logger.Info(fmt.Sprintf("Registering application server %s", appAdress))
	logger.Error(http.ListenAndServe(appAdress, router).Error())
}

func getDbClient() *sqlx.DB {
	dbURL := config.GetDBAddress()
	client, err := sqlx.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
