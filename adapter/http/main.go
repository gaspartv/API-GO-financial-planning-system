package http

import (
	"net/http"

	"github.com/gaspartv/financial-planning-system/adapter/http/actuator"
	"github.com/gaspartv/financial-planning-system/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init is the initialization
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
