package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/postgres"

	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/http/middleware"
	"github.com/pedro-scarelli/wheredidmymoneygo/di"
)

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	accountService, accountUseCase := di.ConfigAccountDI(conn)
	authenticationService := di.ConfigAuthenticationDI(conn)

	router := mux.NewRouter()
	router.Handle("/login", http.HandlerFunc(authenticationService.Login)).Methods("POST")
	router.Handle("/account", http.HandlerFunc(accountService.Create)).Methods("POST")

	protectedRouter := router.PathPrefix("").Subrouter()
	protectedRouter.Use(middleware.JwtAuthorizer(accountUseCase))

	protectedRouter.Handle("/account", http.HandlerFunc(accountService.Update)).Methods("PATCH")
	protectedRouter.Handle("/account/{id}", http.HandlerFunc(accountService.GetByID)).Methods("GET")
	protectedRouter.Handle("/account/{id}", http.HandlerFunc(accountService.Delete)).Methods("DELETE")

	protectedRouter.Handle("/account/movement", http.HandlerFunc(accountService.Movement)).Methods("POST")

	port := os.Getenv("API_PORT")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
