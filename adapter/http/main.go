package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/postgres"
	"github.com/pedro-scarelli/wheredidmymoneygo/di"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	accountService := di.ConfigAccountDI(conn)

	router := mux.NewRouter()
	router.Handle("/account", http.HandlerFunc(accountService.Create)).Methods("POST")
	router.Handle("/account", http.HandlerFunc(accountService.Update)).Methods("PATCH")
	router.Handle("/account", http.HandlerFunc(accountService.Get)).Methods("GET")
	router.Handle("/account/{id}", http.HandlerFunc(accountService.GetByID)).Methods("GET")
	router.Handle("/account/{id}", http.HandlerFunc(accountService.Delete)).Methods("DELETE")

	port := os.Getenv("API_PORT")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
