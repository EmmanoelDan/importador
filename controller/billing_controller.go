package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/EmmanoelDan/importador/service"
	"github.com/gin-gonic/gin"
)

type BillingController struct {
	BillingService *service.BillingService
}

func NewBillingController(billingService *service.BillingService) *BillingController {
	return &BillingController{BillingService: billingService}
}

func (c *BillingController) FindAllWithRelations(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	billings, err := c.BillingService.FindAllWithRelations(page, pageSize)
	if err != nil {
		log.Printf("Error getting all billings: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"billings": billings,
	})
}
