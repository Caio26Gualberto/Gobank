package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

func Connect() {
	connString := "Data Source=localhost,1433;Initial Catalog=Gobank;Trusted_Connection=True;TrustServerCertificate=True;"

	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Error creating connection pool: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %s", err)
	}

	fmt.Println("Successfully connected to SQL Server!")
}

func GetDB() *sql.DB {
	return db
}
