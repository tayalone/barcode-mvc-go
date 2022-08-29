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

	orderInst := courierorder.GetTableStruct(input.CourierCode, input.IsCod)

	orderInst.SetBarcode("")

	r := db.Create(orderInst)

	if r.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Create Order Fail",
		})
	}

	barcode, err := barcode.Gen(orderInst.GetID(), input.CourierCode, input.IsCod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Create Barcode Fail",
		})
	}
	u := db.Model(orderInst).Where("id = ?", orderInst.GetID()).Update("barcode", barcode)

	if u.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Update Order Barcode Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "OK",
		"courierCode": input.CourierCode,
		"isCod":       input.IsCod,
		"order":       orderInst,
	})
}
