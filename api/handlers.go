package api

import (
	"zocket-api/database"

	"github.com/gin-gonic/gin"
)

// handler to handle product creation
func CreateProduct(c *gin.Context) {
	var product database.Product

	// Bind JSON body to the product struct
	if err := c.BindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": "Invalid product data"})
		return
	}

	// Call function to save product to the database
	if err := database.CreateProduct(product); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Product created successfully",
	})
}
