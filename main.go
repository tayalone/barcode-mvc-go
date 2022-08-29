package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	barcodeController "github.com/tayalone/barcode-mvc-go/controller/barcode"
	orderController "github.com/tayalone/barcode-mvc-go/controller/order"
	"github.com/tayalone/barcode-mvc-go/model/rdb"
)

func main() {
	myRdb := rdb.Connect()

	if myRdb.GetStatus() && os.Getenv("RDM_MIGRATION") == "true" {
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

	r.GET("/error", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "OK",
		})
		return
	})

	order := r.Group("/order")
	{
		order.POST("/", orderController.Create)
	}
	barcode := r.Group("/barcode")
	{
		barcode.GET("/", barcodeController.GetAll)
		barcode.GET("/:id", barcodeController.GetByID)
		barcode.POST("/", barcodeController.Create)
		barcode.PATCH("/:id", barcodeController.UpdateByID)
		barcode.DELETE("/:id", barcodeController.DeleteByID)

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
