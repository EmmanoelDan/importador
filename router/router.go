package router

import (
	"log"

	"github.com/EmmanoelDan/importador/controller"
	"github.com/EmmanoelDan/importador/repository"
	"github.com/EmmanoelDan/importador/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	r := gin.Default()

	initializeRouter(r)

	r.Run()
}

func initializeRouter(r *gin.Engine) {
	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	partnerRepo := &repository.PartnerRepository{DB: db}
	customerRepo := &repository.CustomerRepository{DB: db}
	productRepo := &repository.ProductRepository{DB: db}
	skuRepo := &repository.SkuRepository{DB: db}
	entitlementRepo := &repository.EntitlementRepository{DB: db}
	billingRepo := &repository.BillingRepository{DB: db}

	importService := &service.ImportService{
		DB:              db,
		PartnerRepo:     partnerRepo,
		CustomerRepo:    customerRepo,
		ProductRepo:     productRepo,
		SkuRepo:         skuRepo,
		EntitlementRepo: entitlementRepo,
		BillingRepo:     billingRepo,
	}

	importController := controller.NewImportController(importService)

	v1 := r.Group("/api/v1")

	{
		v1.POST("/import_file", importController.UploadCSVHandler)
		v1.GET("/customer")
		v1.GET("/partner")
		v1.GET("/product")
		v1.GET("/sku")
		v1.GET("/entitlement")
		v1.GET("/billing")
	}
}
