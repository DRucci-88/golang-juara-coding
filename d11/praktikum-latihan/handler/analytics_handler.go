package handler

import (
	"net/http"
	"praktikum/service"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandler interface {
	PopularProduct(c *gin.Context)
}

type analyticsHandlerImpl struct {
	orderItemService service.OrderItemService
}

func NewAnaliticsHandler(
	orderItemService service.OrderItemService,
) AnalyticsHandler {
	return &analyticsHandlerImpl{
		orderItemService: orderItemService,
	}
}

func (h *analyticsHandlerImpl) PopularProduct(c *gin.Context) {
	popular, err := h.orderItemService.AnalyticsPopularProduct(0)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": popular})
}
