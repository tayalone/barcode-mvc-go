package barcode

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/barcode-mvc-go/model/rdb"
	"github.com/tayalone/barcode-mvc-go/model/rdb/courierorder"
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
	myRdb, _ := rdb.GetDbInstance()
	db := myRdb.GetDb()

	var bcs []rdb.BarcodeCondition

	db.Order("created_at desc,id desc").Find(&bcs)

	c.JSON(http.StatusOK, gin.H{
		"message":           "OK",
		"barCodeConditions": bcs,
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

	myRdb, _ := rdb.GetDbInstance()
	db := myRdb.GetDb()

	bc := &rdb.BarcodeCondition{}
	r := db.First(bc, gi.ID)

	if r.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":          "OK",
			"barCodeCondition": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "OK",
		"barCodeCondition": bc,
	})
	return
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

	myRdb, _ := rdb.GetDbInstance()
	db := myRdb.GetDb()

	lastBc := &rdb.BarcodeCondition{}
	resLasCond := db.Where(
		&rdb.BarcodeCondition{
			CourierCode: input.CourierCode,
			IsCod:       input.IsCod,
		},
	).First(lastBc)

	if resLasCond.RowsAffected != 1 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Previous Condition Not Found",
		})
		return
	}

	newBc := &rdb.BarcodeCondition{
		CourierCode:   input.CourierCode,
		IsCod:         input.IsCod,
		StartBarcode:  input.StartBarcode,
		BatchSize:     input.BatchSize,
		PrevCondLogID: lastBc.CondLogID,
		CondLogID:     lastBc.CondLogID + uint(input.BatchSize),
	}

	resCreate := db.Create(newBc)

	if resCreate.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Previous Inser Condition Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "OK",
		"barCodeCondition": newBc,
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

	myRdb, _ := rdb.GetDbInstance()
	db := myRdb.GetDb()

	tableName := courierorder.GetTableName(update.CourierCode, update.IsCod)

	var currentID int64
	db.Table(tableName).Count(&currentID)

	bc := &rdb.BarcodeCondition{}
	r := db.First(bc, gi.ID)
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Condition Not Found",
		})
		return
	}

	newCondLogID := bc.PrevCondLogID + uint(update.BatchSize) - 1

	if uint(currentID) >= newCondLogID {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Batch Size Must Bigger than %d", update.BatchSize),
		})
		return
	}

	bc.BatchSize = update.BatchSize
	bc.CondLogID = bc.PrevCondLogID + uint(update.BatchSize)

	db.Save(bc)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
	return
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

	myRdb, _ := rdb.GetDbInstance()
	db := myRdb.GetDb()
	bc := &rdb.BarcodeCondition{}
	r := db.First(bc, gi.ID)
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Condition Not Found",
		})
		return
	}

	tableName := courierorder.GetTableName(bc.CourierCode, bc.IsCod)

	var currentID int64
	db.Table(tableName).Count(&currentID)

	if currentID >= int64(bc.PrevCondLogID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can not Remove used Barcode",
		})
		return
	}

	db.Delete(&rdb.BarcodeCondition{}, gi.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
