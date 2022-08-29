package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderInput struct {
	CourierCode string `json:"courierCode" binding:"required"`
	IsCod       bool   `json:"isCod" binding:"required"`
}

/*
Create is Create New Barcode
*/
func Create(c *gin.Context) {
	var input orderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"todo":    "Create New Order Condition",
		"input":   input,
	})
}
