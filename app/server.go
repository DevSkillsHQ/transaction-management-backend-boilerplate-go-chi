// +build ignore

package main

import (
	"fmt"
	"log"
	"net/http"

	"account-management/db"

	"github.com/go-chi/chi"
)

type balanceResponse struct {
	Balance int `json:"balance"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func submitTransaction(w http.ResponseWriter, r *http.Request) {

	// TODO Implement submitting a new transaction

	w.WriteHeader(201)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {

	// TODO Implement fetching historical transactions

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func getAccount(w http.ResponseWriter, r *http.Request) {

	// Read `accountId` from the request
	//accountId := chi.URLParam(r, "accountId")

	// TODO Implement retrieving account details

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func getTransaction(w http.ResponseWriter, r *http.Request) {

	// Read `transactionId` from the request
	//transactionId := chi.URLParam(r, "transactionId")

	// TODO Implement retrieving transaction

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

// Webserver port
var port = 8080

func main() {

	// Reset and open DB
	dbs := db.Init("db.sqlite", true)
	defer dbs.Close()

	// Define router
	r := chi.NewRouter()
	r.Get("/ping", pingHandler)
	r.Post("/transactions", submitTransaction)
	r.Get("/transactions", getTransactions)
	r.Get("/accounts/{accountId}", getAccount)
	r.Get("/transactions/{transactionId}", getTransaction)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
	fmt.Println("Serving on port ", port)
}
