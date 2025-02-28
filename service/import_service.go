package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/EmmanoelDan/importador/model"
	"github.com/EmmanoelDan/importador/repository"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ImportService struct {
	PartnerRepo *repository.PartnerRepository
	CustomerRepo *repository.CustomerRepository
	ProductRepo *repository.ProductRepository
	SkuRepo *repository.SkuRepository
	Subscription *repository.SubscriptionRepository
	Meter *repository.MeterRepository
	Billing *repository.Billing
	DB *gorm.DB
}
const workerPoolSize = 1000

func (s *ImportService) ImportCSV(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return fmt.Errorf("error opening CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
        return fmt.Errorf("error reading CSV records: %w", err)
    }

	if len(records) < 2 {
		return fmt.Errorf("CSV file is empty or only contains headers")
    }

	headers := records[0]
	rows := records[1:]

	rowChan := make(chan map[string]string, len(rows))
	var waitGroup sync.WaitGroup
	for i := 0; i < workerPoolSize; i++ {
		waitGroup.Add(1)
		go func ()  {
			defer waitGroup.Done()
			for rowMap := range rowChan {
				s.processRow(rowMap)
			}
		}()
	}

	for _, record := range rows {
		rowMap := make(map[string]string, len(headers))
		for i, header := range headers {
            rowMap[header] = record[i]
        }
		rowChan <- rowMap
	}

	close(rowChan)
	waitGroup.Wait()

	return nil
}

func (s *ImportService) processRow(rowMap map[string]string) {
	tx := s.DB.Begin()

	defer func(){
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Recovery from panic: %v", r)
		}
	}()

	partner := &model.Partner{
        PartnerId:   rowMap["PartnerId"],
        PartnerName: rowMap["PartnerName"],
		MpnId: rowMap["MpnId"],
        Tier2MpnId: rowMap["Tier2MpnId"],
    }
	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "partnerid"}},
		DoNothing: true,
	}).Create(partner).Error; err != nil {
		log.Printf("Error inserting partner: %v\n", err)
		tx.Rollback()
		return
	}

	customer := &model.Customer{
		CustomerId: rowMap["CustomerId"],
        CustomerName: rowMap["CustomerName"],
        CustomerDomainName: rowMap["CustomerDomainName"],
        CustomerCountry: rowMap["CustomerCountry"],
		PartnerId: partner.PartnerId,
	}

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "customerid"}},
        DoNothing: true,
	}).Create(customer).Error; err != nil {
		log.Printf("Error inserting customer: %v\n", err)
        tx.Rollback()
        return
	}

	product := &model.Product{
		ProductId: rowMap["ProductId"],
        ProductName: rowMap["ProductName"],
        PublisherName: rowMap["PublisherName"],
        PublisherId: rowMap["PublisherId"],
	}

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "productid"}},
        DoNothing: true,
	}).Create(product).Error; err != nil {
		log.Printf("Error inserting product: %v\n", err)
        tx.Rollback()
        return
	}

	sku := &model.Sku{
		ProductId: rowMap["ProductId"],
        SkuId: rowMap["SkuId"],
        SkuName: rowMap["SkuName"],
        AvailabilityId: rowMap["AvailabilityId"],
	}

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "skuid"}},
        DoNothing: true,
	}).Create(sku).Error; err != nil {
		log.Printf("Error inserting SKU: %v\n", err)
        tx.Rollback()
        return
	}

	subscription := &model.Subscription{
		SubscriptionId: rowMap["SubscriptionId"],
		CustomerId: customer.CustomerId,
        ProductId: product.ProductId,
        SubscriptionDescription: rowMap["SubscriptionDescription"],
	}

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "subscriptionid"}},
        DoNothing: true,
	}).Create(subscription).Error; err != nil {
		log.Printf("Error inserting subscription: %v\n", err)
        tx.Rollback()
        return
	}

	meter := &model.Meter{
		MeterId: rowMap["MeterId"],
        MeterType: rowMap["MeterType"],
		MeterCategory: rowMap["MeterCategory"],
		MeterSubCategory: rowMap["MeterSubCategory"],
		MeterName: rowMap["MeterName"],
		MeterRegion: rowMap["MeterRegion"],
		Unit: rowMap["Unit"],
	}

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "meterid"}},
        DoNothing: true,
	}).Create(meter).Error; err != nil {
		log.Printf("Error inserting meter: %v\n", err)
        tx.Rollback()
        return
	}


	entitlement := &model.Entitlement{
		EntitlementId: rowMap["EntitlementId"],
		EntitlementDescription: rowMap["EntitlementDescription"],
		PartnerEarnedCreditPercentage: rowMap["PartnerEarnedCreditPercentage"],
		CreditPercentage: rowMap["CreditPercentage"],
		CreditType: rowMap["CreditType"],
		BenefitOrderId: rowMap["BenefitOrderId"],
		BenefitId: rowMap["BenefitId"],
		BenefitType: rowMap["BenefitType"],
	}

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "entitlementid"}},
        DoNothing: true,
	}).Create(entitlement).Error; err != nil {
		log.Printf("Error inserting entitlement: %v\n", err)
        tx.Rollback()
        return
	}

	billing := &model.Billing{
		PartnerId: partner.PartnerId,
		SubscriptionId : subscription.SubscriptionId,
		MeterId: meter.MeterId,
		EntitlementId: entitlement.EntitlementId,
		InvoiceNumber: rowMap["InvoiceNumber"],
		ChargeStartDate: rowMap["ChargeStartDate"],
		ChargeEndDate: rowMap["ChargeEndDate"],
		UsageDate: rowMap["UsageDate"],
		ResourceLocation: rowMap["ResourceLocation"],
		ConsumedService: rowMap["ConsumedService"],
		ResourceGroup: rowMap["ResourceGroup"],
		ResourceURI: rowMap["ResourceURI"],
		ChargeType: rowMap["ChargeType"],
		UnitPrice: rowMap["UnitPrice"],
		Quantity: rowMap["Quantity"],
		UnitType: rowMap["UnitType"],
		BillingPreTaxTotal: rowMap["BillingPreTaxTotal"],
		BillingCurrency: rowMap["BillingCurrency"],
		PricingPreTaxTotal: rowMap["PricingPreTaxTotal"],
		PricingCurrency: rowMap["PricingCurrency"],
		ServiceInfo1: rowMap["ServiceInfo1"],
		ServiceInfo2: rowMap["ServiceInfo2"],
		Tags: datatypes.JSON(rowMap["Tags"]),
		AdditionalInfo: datatypes.JSON(rowMap["AdditionalInfo"]),
		EffectiveUnitPrice: rowMap["EffectiveUnitPrice"],
		PCToBCExchangeRate: rowMap["PCToBCExchangeRate"],
		PCToBCExchangeRateDate: rowMap["PCToBCExchangeRateDate"],
	}

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "billingid"}},
        DoNothing: true,
	}).Create(billing).Error; err != nil {
		log.Printf("Error inserting billing: %v\n", err)
        tx.Rollback()
        return
	}

	tx.Commit()
}

