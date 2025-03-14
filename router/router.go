package router

import (
	"fmt"
	"log"
	"os"

	"github.com/EmmanoelDan/importador/controller"
	"github.com/EmmanoelDan/importador/middleware"
	"github.com/EmmanoelDan/importador/repository"
	"github.com/EmmanoelDan/importador/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	r := gin.Default()
	initializeRouter(r)

	port := ":8080"
	log.Printf("Server is listening on port %s", port)

	r.Run(port)
}

func initializeRouter(r *gin.Engine) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
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

	userRepo := &repository.UserRepository{DB: db}

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
	userRegisterService := service.NewCreateUserService(userRepo)
	userRegisterController := controller.NewUserHandler(userRegisterService)
	authUserService := service.NewAuthService(*userRepo)
	authController := controller.NewAuthHandler(authUserService)
	billingService := service.NewBillingService(billingRepo)
	billingController := controller.NewBillingController(billingService)

	r.POST("/sign", authController.Sign)
	r.POST("/register", userRegisterController.Register)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	{
		protected.POST("/import_file", importController.UploadCSVHandler)
		protected.GET("/billings", billingController.FindAllWithRelations)
	}
}
