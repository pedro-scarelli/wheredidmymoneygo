package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/postgres"
	"github.com/pedro-scarelli/wheredidmymoneygo/di"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	accountService := di.ConfigAccountDI(conn)

	router := mux.NewRouter()
	router.Handle("/account", http.HandlerFunc(accountService.Create)).Methods("POST")
	router.Handle("/account", http.HandlerFunc(accountService.Get)).Methods("GET")
	router.Handle("/account/{id}", http.HandlerFunc(accountService.GetByID)).Methods("GET")
	router.Handle("/account/{id}", http.HandlerFunc(accountService.Update)).Methods("PATCH")
	router.Handle("/account/{id}", http.HandlerFunc(accountService.Delete)).Methods("DELETE")

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
