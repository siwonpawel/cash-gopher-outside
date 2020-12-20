package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/siwonpawel/cash-gopher-outside/banking-auth/config"
	"github.com/siwonpawel/cash-gopher-outside/banking-auth/domain"
	"github.com/siwonpawel/cash-gopher-outside/banking-auth/service"
)

func Start() {
	router := mux.NewRouter()
	authRepository := domain.NewAuthRepository(getDbClient())
	ah := AuthHandler{service.NewLoginService(authRepository, domain.GetRolePermissions())}

	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	appAddress := config.GetAppAddress()
	log.Println(fmt.Sprintf("Starting OAuth server on %s", appAddress))
	log.Fatal(http.ListenAndServe(appAddress, router))
}

func getDbClient() *sqlx.DB {

	client, err := sqlx.Open("mysql", config.GetDBAddress())
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
