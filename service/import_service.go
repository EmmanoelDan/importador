package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/EmmanoelDan/importador/model"
	"github.com/EmmanoelDan/importador/repository"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ImportService struct {
	PartnerRepo     *repository.PartnerRepository
	CustomerRepo    *repository.CustomerRepository
	ProductRepo     *repository.ProductRepository
	SkuRepo         *repository.SkuRepository
	EntitlementRepo *repository.EntitlementRepository
	BillingRepo     *repository.BillingRepository
	DB              *gorm.DB
}

const (
	workerPoolSize = 6
	batchSize      = 100
)

func (s *ImportService) ImportCSV(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return fmt.Errorf("error opening CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("error reading CSV headers: %w", err)
	}

	rowChan := make(chan map[string]string, batchSize*workerPoolSize)
	var wg sync.WaitGroup

	for i := 0; i < workerPoolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.processBatch(rowChan)
		}()
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		rowMap := make(map[string]string, len(headers))
		for i, header := range headers {
			rowMap[header] = record[i]
		}

		rowChan <- rowMap
	}

	close(rowChan)
	wg.Wait()

	return nil
}

func (s *ImportService) processBatch(rowChan <-chan map[string]string) {
	var partners []model.Partner
	var customers []model.Customer
	var products []model.Product
	var skus []model.Sku
	var entitlements []model.Entitlement
	var billings []model.Billing

	for rowMap := range rowChan {
		partners = append(partners, model.Partner{
			PartnerId:   rowMap["PartnerId"],
			PartnerName: rowMap["PartnerName"],
			MpnId:       rowMap["MpnId"],
			Tier2MpnId:  rowMap["Tier2MpnId"],
		})

		customers = append(customers, model.Customer{
			CustomerId:         rowMap["CustomerId"],
			CustomerName:       rowMap["CustomerName"],
			CustomerDomainName: rowMap["CustomerDomainName"],
			CustomerCountry:    rowMap["CustomerCountry"],
			PartnerId:          partners[len(partners)-1].PartnerId,
		})

		skus = append(skus, model.Sku{
			SkuId:          rowMap["SkuId"],
			SkuName:        rowMap["SkuName"],
			AvailabilityId: rowMap["AvailabilityId"],
		})

		products = append(products, model.Product{
			ProductId:               rowMap["ProductId"],
			SkuId:                   skus[len(skus)-1].SkuId,
			ProductName:             rowMap["ProductName"],
			PublisherName:           rowMap["PublisherName"],
			PublisherId:             rowMap["PublisherId"],
			SubscriptionId:          rowMap["SubscriptionId"],
			SubscriptionDescription: rowMap["SubscriptionDescription"],
		})

		partnerEarnedCreditPercentage, err := strconv.ParseFloat(rowMap["PartnerEarnedCreditPercentage"], 64)
		if err != nil {
			log.Fatal("Error parsing PartnerEarnedCreditPercentage:", err)
		}

		creditPercentage, err := strconv.ParseFloat(rowMap["CreditPercentage"], 64)
		if err != nil {
			log.Fatal("Error parsing CreditPercentage:", err)
		}

		entitlements = append(entitlements, model.Entitlement{
			EntitlementId:                 rowMap["EntitlementId"],
			EntitlementDescription:        rowMap["EntitlementDescription"],
			PartnerEarnedCreditPercentage: partnerEarnedCreditPercentage,
			CreditPercentage:              creditPercentage,
			CreditType:                    rowMap["CreditType"],
			BenefitOrderId:                rowMap["BenefitOrderId"],
			BenefitId:                     rowMap["BenefitId"],
			BenefitType:                   rowMap["BenefitType"],
		})

		layout := "1/2/2006"

		chargeStartDate, err := time.Parse(layout, rowMap["ChargeStartDate"])
		if err != nil {
			log.Fatal("Error parsing ChargeStartDate:", err)
		}
		chargeEndDate, err := time.Parse(layout, rowMap["ChargeEndDate"])
		if err != nil {
			log.Fatal("Error parsing ChargeEndDate:", err)
		}
		usageDate, err := time.Parse(layout, rowMap["UsageDate"])
		if err != nil {
			log.Fatal("Error parsing UsageDate:", err)
		}
		pCToBCExchangeRateDate, err := time.Parse(layout, rowMap["PCToBCExchangeRateDate"])
		if err != nil {
			log.Fatal("Error parsing PCToBCExchangeRateDate:", err)
		}

		unitPrice, err := strconv.ParseFloat(rowMap["UnitPrice"], 64)
		if err != nil {
			log.Fatal("Error parsing UnitPrice:", err)
		}
		quantity, err := strconv.ParseFloat(rowMap["Quantity"], 64)
		if err != nil {
			log.Fatal("Error parsing Quantity:", err)
		}
		billingPreTaxTotal, err := strconv.ParseFloat(rowMap["BillingPreTaxTotal"], 64)
		if err != nil {
			log.Fatal("Error parsing BillingPreTaxTotal:", err)
		}
		pricingPreTaxTotal, err := strconv.ParseFloat(rowMap["PricingPreTaxTotal"], 64)
		if err != nil {
			log.Fatal("Error parsing PricingPreTaxTotal:", err)
		}
		effectiveUnitPrice, err := strconv.ParseFloat(rowMap["EffectiveUnitPrice"], 64)
		if err != nil {
			log.Fatal("Error parsing EffectiveUnitPrice:", err)
		}
		pCToBCExchangeRate, err := strconv.ParseFloat(rowMap["PCToBCExchangeRate"], 64)
		if err != nil {
			log.Fatal("Error parsing PCToBCExchangeRate:", err)
		}

		billings = append(billings, model.Billing{
			CustomerId:             customers[len(customers)-1].CustomerId,
			ProductId:              products[len(products)-1].ProductId,
			EntitlementId:          entitlements[len(entitlements)-1].EntitlementId,
			InvoiceNumber:          rowMap["InvoiceNumber"],
			ChargeStartDate:        chargeStartDate,
			ChargeEndDate:          chargeEndDate,
			UsageDate:              usageDate,
			MeterId:                rowMap["MeterId"],
			MeterType:              rowMap["MeterType"],
			MeterCategory:          rowMap["MeterCategory"],
			MeterSubCategory:       rowMap["MeterSubCategory"],
			MeterName:              rowMap["MeterName"],
			MeterRegion:            rowMap["MeterRegion"],
			Unit:                   rowMap["Unit"],
			ResourceLocation:       rowMap["ResourceLocation"],
			ConsumedService:        rowMap["ConsumedService"],
			ResourceGroup:          rowMap["ResourceGroup"],
			ResourceURI:            rowMap["ResourceURI"],
			ChargeType:             rowMap["ChargeType"],
			UnitPrice:              unitPrice,
			Quantity:               quantity,
			UnitType:               rowMap["UnitType"],
			BillingPreTaxTotal:     billingPreTaxTotal,
			BillingCurrency:        rowMap["BillingCurrency"],
			PricingPreTaxTotal:     pricingPreTaxTotal,
			PricingCurrency:        rowMap["PricingCurrency"],
			ServiceInfo1:           rowMap["ServiceInfo1"],
			ServiceInfo2:           rowMap["ServiceInfo2"],
			Tags:                   datatypes.JSON(rowMap["Tags"]),
			AdditionalInfo:         datatypes.JSON(rowMap["AdditionalInfo"]),
			EffectiveUnitPrice:     effectiveUnitPrice,
			PCToBCExchangeRate:     pCToBCExchangeRate,
			PCToBCExchangeRateDate: pCToBCExchangeRateDate,
		})

		if len(partners) >= batchSize {
			s.insertBatch(partners, customers, skus, products, entitlements, billings)
			partners = nil
			customers = nil
			skus = nil
			products = nil
			entitlements = nil
			billings = nil
		}
	}

	if len(partners) > 0 {
		s.insertBatch(partners, customers, skus, products, entitlements, billings)
	}
}

func (s *ImportService) insertBatch(partners []model.Partner, customers []model.Customer, skus []model.Sku, products []model.Product, entitlements []model.Entitlement, billings []model.Billing) {
	tx := s.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("panic recovery: %v", r)
		}
	}()

	if err := tx.Clauses(clause.OnConflict{Columns: []clause.Column{{
		Name: "partnerid"}},
		DoNothing: true,
	}).Create(&partners).Error; err != nil {
		log.Printf("Erro ao inserir batch de partners: %v", err)
		tx.Rollback()
		return
	}

	if err := tx.Clauses(clause.OnConflict{Columns: []clause.Column{{
		Name: "customerid"}},
		DoNothing: true,
	}).Create(&customers).Error; err != nil {
		log.Printf("Erro ao inserir batch de customers: %v", err)
		tx.Rollback()
		return
	}

	if err := tx.Clauses(clause.OnConflict{Columns: []clause.Column{{
		Name: "skuid"}},
		DoNothing: true,
	}).Create(&skus).Error; err != nil {
		log.Printf("Erro ao inserir batch de skus: %v", err)
		tx.Rollback()
		return
	}

	if err := tx.Clauses(clause.OnConflict{Columns: []clause.Column{{
		Name: "productid"}},
		DoNothing: true,
	}).Create(&products).Error; err != nil {
		log.Printf("Erro ao inserir batch de products: %v", err)
		tx.Rollback()
		return
	}

	if err := tx.Clauses(clause.OnConflict{Columns: []clause.Column{{
		Name: "entitlementid"}},
		DoNothing: true,
	}).Create(&entitlements).Error; err != nil {
		log.Printf("Erro ao inserir batch de meters: %v", err)
		tx.Rollback()
		return
	}

	if err := tx.Create(&billings).Error; err != nil {
		log.Printf("Erro ao inserir batch de billing: %v", err)
		tx.Rollback()
		return
	}

	tx.Commit()
}
