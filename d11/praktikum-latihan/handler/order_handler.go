package handler

import (
	"log"
	"net/http"
	"praktikum/dto"
	"praktikum/service"

	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	Create(c *gin.Context)
}

type orderHandlerImpl struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) OrderHandler {
	return &orderHandlerImpl{
		orderService: orderService,
	}
}

func (h *orderHandlerImpl) Create(c *gin.Context) {
	var input dto.CheckoutRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	log.Printf("%+v", input)
	order, err := h.orderService.Create(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Transaksi checkout berhasil di proses",
		"data":    order,
	})

}
