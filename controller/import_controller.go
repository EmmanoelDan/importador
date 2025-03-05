package controller

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/EmmanoelDan/importador/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ImportController struct {
	ImportService *service.ImportService
}

func NewImportController(importService *service.ImportService) *ImportController {
	return &ImportController{ImportService: importService}
}

func (c *ImportController) UploadCSVHandler(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ext := filepath.Ext(file.Filename)
	randomName := uuid.New().String() + ext
	filePath := "./temp/" + randomName
	if err := os.MkdirAll("./temp", os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := c.ImportService.ImportCSV(filePath); err != nil {
			log.Printf("Error importing service: %s", err)
		}

		if err := os.Remove(filePath); err != nil {
			log.Printf("Error ao remover arquivo temporario: %s", err)
		}

	}()

	wg.Wait()

	ctx.JSON(http.StatusOK, gin.H{"message": "Upload successfully imported"})
}
