package main

import (
	"fmt"
	"log"

	"github.com/EmmanoelDan/importador/repository"
	"github.com/EmmanoelDan/importador/service"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	dbUser = "postgres"
// 	dbPass = "root"
// 	dbName = "db_import"
// 	dbHost = "localhost"
// 	dbPort = "5432"
// )
func main() {
	dsn := "host=localhost user=postgres password=root dbname=db_import port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	partnerRepo := &repository.PartnerRepository{DB: db}

	importService := &service.ImportService{
		PartnerRepo: partnerRepo,
	}

	if err := importService.ImportCSV("Reconfile-fornecedores.csv"); err != nil {
		log.Fatal("Error importing CSV: ", err)
	}

	fmt.Println("CSV data imported successfully!")

}