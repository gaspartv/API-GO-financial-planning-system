package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gaspartv/financial-planning-system/model/transaction"
	"github.com/gaspartv/financial-planning-system/util"
)

// GetTransactions returns a list of transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	salaryReceived := util.StringToTime("2019-02-12T10:10:10")

	transactions := transaction.Transactions{
		transaction.Transaction{
			Title: "Salário",
			Amount: 1200.0,
			Type: 0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

// CreateTransaction creates a new transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	res := transaction.Transactions{}
	body, _ := ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
