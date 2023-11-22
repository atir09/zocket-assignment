package api

import (
	"github.com/gin-gonic/gin"
)

// handler to handle product creation
func CreateProduct(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Product created successfully",
	})
}
