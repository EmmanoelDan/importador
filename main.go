package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbUser = "postgres"
	dbPass = "root"
	dbName = "db_import"
	dbHost = "localhost"
	dbPort = "5432"
)
func main() {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPass, dbName, dbHost, dbPort)
	
	db, err := sql.Open("postgres", connStr)
	
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()
	
	fmt.Println("Connecting to database successfully")

	file, err := os.Open("Reconfile-fornecedores.csv")
	if err != nil {
        log.Fatal("Error open CSV file: ", err)
    }
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
        log.Fatal("Error reading CSV file: ", err)
    }

	for _, row := range records {
		fmt.Println(row)
	}

}