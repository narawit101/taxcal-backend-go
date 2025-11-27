package handlers

import (
	"net/http"
	"test-backend/models"
	"test-backend/services"
	"test-backend/utils"

	"github.com/gin-gonic/gin"
)

type TaxResponseOutput struct {
	Tax float64 `json:"tax"`
}

func CalculateTaxHandlerSimple(c *gin.Context) {
	var req models.TaxRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	if err := utils.ValidateTaxRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := services.CalculateTax(req)
	output := TaxResponseOutput{
		Tax: result.Tax,
	}

	c.JSON(http.StatusOK, output)
}

func CalculateTaxHandlerDetail(c *gin.Context) {
	var req models.TaxRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	if err := utils.ValidateTaxRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := services.CalculateTax(req)

	c.JSON(http.StatusOK, result)
}
