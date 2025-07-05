package main

import (
	"log"
	"net/http"

	"github.com/Caio26Gualberto/gobank/internal/account"
	accRepo "github.com/Caio26Gualberto/gobank/internal/account/repository"
	"github.com/Caio26Gualberto/gobank/internal/api"
	"github.com/Caio26Gualberto/gobank/internal/db"
	"github.com/Caio26Gualberto/gobank/internal/transaction"
	transRepo "github.com/Caio26Gualberto/gobank/internal/transaction/repository"
)

func main() {
	db.Connect()
	database := db.GetDB()
	log.Println("DB connection initialized.")

	accountRepo := accRepo.NewSQLAccountRepository(database)
	accountHandler := account.NewAccountHandler(accountRepo)

	transactionRepo := transRepo.NewSQLTransactionRepository(database)
	transactionHandler := transaction.NewTransactionHandler(transactionRepo)

	r := api.InitRouters(accountHandler, transactionHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
