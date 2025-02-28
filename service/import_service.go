package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/EmmanoelDan/importador/model"
	"github.com/EmmanoelDan/importador/repository"
	"gorm.io/gorm/clause"
)

type ImportService struct {
	PartnerRepo *repository.PartnerRepository
}

func (s *ImportService) ImportCSV(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return fmt.Errorf("error opening CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()

	if err != nil {
        return fmt.Errorf("error reading CSV records: %w", err)
    }
	fmt.Println("CSV Headers:", headers)

	records, err := reader.ReadAll()

	if err != nil {
        return fmt.Errorf("error reading CSV records: %w", err)
    }

	for _, row := range records {

		partner := &model.Partner{
            PartnerId: row[0],
			PartnerName: row[1],
        }

		result := s.PartnerRepo.DB.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "partnerid"}},
			DoNothing: true,
		}).Create(partner)

		if result.Error != nil {
			log.Printf("Error inserting partner: %w", result.Error)
		}
	}

	return nil

}