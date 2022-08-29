package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/barcode-mvc-go/model/rdb"
)

func main() {
	myRdb := rdb.Connect()

	if myRdb.GetStatus() {
		myRdb.AutoMigrate()
	}

	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		_, errRDB := rdb.GetDbInstance()
		if errRDB != nil {
			c.JSON(http.StatusTeapot, gin.H{
				"message": errRDB.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
		return
	})

	order := r.Group("/")
	{
		type orderInput struct {
			CourierCode  string `json:"courierCode" binding:"required"`
			IsCod        bool   `json:"isCod" binding:"required"`
			StartBarcode string `json:"startBarcode" binding:"required"`
			BatchSize    uint32 `json:"batchSize" binding:"required"`
		}

		order.POST("/", func(c *gin.Context) {
			var input orderInput
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"todo":    "Create New Barcode ",
				"input":   input,
			})
		})
	}
	barcode := r.Group("/barcode")
	{
		barcode.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"todo":    "Get all condition",
			})
		})

		type getID struct {
			ID uint `uri:"id" binding:"required"`
		}

		barcode.GET("/:id", func(c *gin.Context) {
			var gi getID
			if err := c.ShouldBindUri(&gi); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"todo":    fmt.Sprintf("Get Condition ID: %d", gi.ID),
			})
		})

		type barCodeInput struct {
			CourierCode  string `json:"courierCode" binding:"required"`
			IsCod        bool   `json:"isCod" binding:"required"`
			StartBarcode string `json:"startBarcode" binding:"required"`
			BatchSize    uint32 `json:"batchSize" binding:"required"`
		}

		barcode.POST("/", func(c *gin.Context) {
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
		})

		type barCodeUpdate struct {
			CourierCode string `json:"courierCode" binding:"required"`
			IsCod       bool   `json:"isCod" binding:"required"`
			BatchSize   uint32 `json:"batchSize" binding:"required"`
		}

		barcode.PATCH("/:id", func(c *gin.Context) {
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
		})

		barcode.DELETE("/:id", func(c *gin.Context) {
			var gi getID
			if err := c.ShouldBindUri(&gi); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"todo":    fmt.Sprintf("Remove Condition ID: %d", gi.ID),
			})
		})

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
