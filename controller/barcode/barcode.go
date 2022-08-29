package barcode

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getID struct {
	ID uint `uri:"id" binding:"required"`
}

type barCodeInput struct {
	CourierCode  string `json:"courierCode" binding:"required"`
	IsCod        bool   `json:"isCod" binding:"required"`
	StartBarcode string `json:"startBarcode" binding:"required"`
	BatchSize    uint32 `json:"batchSize" binding:"required"`
}

type barCodeUpdate struct {
	CourierCode string `json:"courierCode" binding:"required"`
	IsCod       bool   `json:"isCod" binding:"required"`
	BatchSize   uint32 `json:"batchSize" binding:"required"`
}

/*
GetAll is Get All Barcode Condition
*/
func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"todo":    "Get all condition",
	})
}

/*
GetByID is Get Barcode By Id Condition
*/
func GetByID(c *gin.Context) {
	var gi getID
	if err := c.ShouldBindUri(&gi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"todo":    fmt.Sprintf("Get Condition ID: %d", gi.ID),
	})
}

/*
Create is Create New Barcode
*/
func Create(c *gin.Context) {
	var input barCodeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"todo":    "Create New Barcode Condition",
		"input":   input,
	})
}

/*
UpdateByID is Barcode Condition By IDs
*/
func UpdateByID(c *gin.Context) {
	var gi getID
	if err := c.ShouldBindUri(&gi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	var update barCodeUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"todo":    fmt.Sprintf("Update Condition ID: %d", gi.ID),
		"update":  update,
	})
}

/*
DeleteByID is Delete Barcode By Id
*/
func DeleteByID(c *gin.Context) {
	var gi getID
	if err := c.ShouldBindUri(&gi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"todo":    fmt.Sprintf("Remove Condition ID: %d", gi.ID),
	})
}
