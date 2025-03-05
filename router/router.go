package router

import (
	"log"

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
	r.POST("/sign", authController.Sign)
	r.POST("/register", userRegisterController.Register)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	{
		protected.POST("/import_file", importController.UploadCSVHandler)
		protected.GET("/billings")
	}
}
