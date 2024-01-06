package http

import (
	"net/http"

	"github.com/gaspartv/financial-planning-system/adapter/http/actuator"
	"github.com/gaspartv/financial-planning-system/adapter/http/transaction"
)

// Init is the initialization
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)
	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
