package api

import (
	"github.com/Caio26Gualberto/gobank/internal/account"
	"github.com/Caio26Gualberto/gobank/internal/transaction"
	"github.com/gorilla/mux"
)

func InitRouters(accountHandler *account.AccountHandler, transactionHandler *transaction.TransactionHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/accounts", accountHandler.CreateAccount()).Methods("POST")
	r.HandleFunc("/accounts/{id}", accountHandler.GetAccount).Methods("GET")
	r.HandleFunc("/getAccounts", accountHandler.GetAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}", accountHandler.UpdateAccount).Methods("PUT")
	r.HandleFunc("/accounts/{id}", accountHandler.DeleteAccount).Methods("DELETE")

	r.HandleFunc("/transactions", transactionHandler.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions/{id}", transactionHandler.GetTransactionsByAccount).Methods("GET")

	return r
}
