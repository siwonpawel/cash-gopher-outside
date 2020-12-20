package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siwonpawel/cash-gopher-outside/banking/dto"
	"github.com/siwonpawel/cash-gopher-outside/banking/logger"
	"github.com/siwonpawel/cash-gopher-outside/banking/service"
)

type AccountHandlers struct {
	service service.DefaultAccountService
}

func (ah AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {

	var nareq dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&nareq)
	if err != nil {
		logger.Error("Error decoding data of new account" + err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error)
	}

	nareq.CustomerID = mux.Vars(r)["customer_id"]

	naresp, appErr := ah.service.NewAccount(nareq)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
	}

	writeResponse(w, http.StatusCreated, naresp)
}
