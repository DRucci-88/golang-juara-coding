package handler

import (
	"errors"
	"log"
	"net/http"
	"praktikum/dto"
	"praktikum/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler interface {
	Create(c *gin.Context)
	FindByID(c *gin.Context)
	Cancel(c *gin.Context)
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

func (h *orderHandlerImpl) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	order, err := h.orderService.FindById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Faktur pesananan tidak ditemukan "})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func (h *orderHandlerImpl) Cancel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	order, err := h.orderService.Cancel(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})

}
