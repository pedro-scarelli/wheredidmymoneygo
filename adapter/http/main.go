package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedro-scarelli/go_login/adapter/postgres"
	"github.com/pedro-scarelli/go_login/di"
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
	userService := di.ConfigUserDI(conn)

	router := mux.NewRouter()
	router.Handle("/user", http.HandlerFunc(userService.Create)).Methods("POST")
	router.Handle("/user", http.HandlerFunc(userService.Get)).Methods("GET")
	router.Handle("/user/{id}", http.HandlerFunc(userService.GetByID)).Methods("GET")
	router.Handle("/user/{id}", http.HandlerFunc(userService.Update)).Methods("PATCH")
	router.Handle("/user/{id}", http.HandlerFunc(userService.Delete)).Methods("DELETE")

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
