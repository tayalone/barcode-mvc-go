package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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

		type getByID struct {
			id uint `uri:"id" binding:"required,id"`
		}
		barcode.GET("/:id", func(c *gin.Context) {
			var getByID getByID
			if err := c.ShouldBindUri(&getByID); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"todo":    fmt.Sprintf("Get Condition ID: %d", getByID.id),
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

		barcode.PATCH("/", func(c *gin.Context) {
			var getByID getByID
			if err := c.ShouldBindUri(&getByID); err != nil {
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
				"todo":    fmt.Sprintf("Update Condition ID: %d", getByID.id),
				"update":  update,
			})
		})

		barcode.DELETE("/:id", func(c *gin.Context) {
			var getByID getByID
			if err := c.ShouldBindUri(&getByID); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"todo":    fmt.Sprintf("Remove Condition ID: %d", getByID.id),
			})
		})

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
