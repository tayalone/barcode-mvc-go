package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("RDM_HOST"),
		os.Getenv("RDM_USER"),
		os.Getenv("RDM_PASSWORD"),
		os.Getenv("RDM_DB"),
		os.Getenv("RDM_PORT"),
		os.Getenv("TIME_ZONE"))

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("wait a minute")

	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

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
