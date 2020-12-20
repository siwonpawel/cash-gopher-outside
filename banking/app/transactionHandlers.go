package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siwonpawel/cash-gopher-outside/banking/dto"
	"github.com/siwonpawel/cash-gopher-outside/banking/errs"
	"github.com/siwonpawel/cash-gopher-outside/banking/service"
)

type TransactionHandlers struct {
	service service.DefaultTransactionService
}

func NewTransactionHandlers(s service.DefaultTransactionService) TransactionHandlers {
	return TransactionHandlers{service: s}
}

func (th TransactionHandlers) AddTransaction(w http.ResponseWriter, r *http.Request) {

	var tr dto.TransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		writeResponse(w, http.StatusBadRequest, errs.NewParseError(err.Error()))
		return
	}

	params := mux.Vars(r)
	tr.CustomerID = params["customer_id"]
	tr.AccountID = params["account_id"]

	response, err := th.service.AddTransaction(tr)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, response)
}
