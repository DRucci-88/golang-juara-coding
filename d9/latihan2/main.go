package main

import (
	"latihan2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var carts = []model.CartItem{
	{ProductID: 1, Quantity: 10},
	{ProductID: 2, Quantity: 40},
	{ProductID: 3, Quantity: 50},
}

func main() {
	r := gin.Default()

	api := r.Group("/api")
	api.POST("/cart-item", func(c *gin.Context) {
		cart := model.CartItem{}

		if err := c.ShouldBindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		carts = append(carts, cart)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Cart Added",
			"data":    cart,
		})
	})

	api.GET("/cart-item", func(c *gin.Context) {
		c.JSON(http.StatusOK, carts)
	})

	api.POST("/checkout", func(c *gin.Context) {
		checkoutRequest := model.CheckoutRequest{}

		if err := c.ShouldBindJSON(&checkoutRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var total float64
		for _, cart := range carts {
			total += float64(cart.Quantity) * 100_000.0
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": gin.H{
				"status": "PENDING_PAYMENT",
				"total":  total,
				"cart":   carts,
			},
		})
	})

	r.Run(":8080")

}
