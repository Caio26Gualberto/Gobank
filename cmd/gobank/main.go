package main

import (
	"log"
	"net/http"

	"github.com/Caio26Gualberto/gobank/internal/account"
	"github.com/Caio26Gualberto/gobank/internal/account/repository"
	"github.com/Caio26Gualberto/gobank/internal/api"
	"github.com/Caio26Gualberto/gobank/internal/db"
)

func main() {
	db.Connect()
	database := db.GetDB()
	log.Println("DB connection initialized.")

	accountRepo := repository.NewSQLAccountRepository(database)
	accountHandler := account.NewAccountHandler(accountRepo)

	r := api.InitRouters(accountHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
