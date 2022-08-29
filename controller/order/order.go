package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/barcode-mvc-go/controller/barcode"
	"github.com/tayalone/barcode-mvc-go/model/rdb"
	"github.com/tayalone/barcode-mvc-go/model/rdb/courierorder"
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

	myRdb, _ := rdb.GetDbInstance()
	db := myRdb.GetDb()

	orderStruct := courierorder.GetTableStruct(input.CourierCode, input.IsCod)

	orderStruct.SetBarcode("")

	r := db.Create(orderStruct)

	if r.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Create Order Fail",
		})
	}

	barcode, err := barcode.Gen(orderStruct.GetID(), input.CourierCode, input.IsCod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Create Barcode Fail",
		})
	}
	u := db.Model(orderStruct).Where("id = ?", orderStruct.GetID()).Update("barcode", barcode)

	if u.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Update Order Barcode Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "OK",
		"courierCode": input.CourierCode,
		"isCod":       input.IsCod,
		"order":       orderStruct,
	})
}
