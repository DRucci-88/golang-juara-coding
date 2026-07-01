package handler

import (
	"net/http"
	"strconv"
	"ujian2_rematch/dto"
	"ujian2_rematch/service"

	"github.com/gin-gonic/gin"
)

type PositionHandler struct {
	positionService *service.PositionService
}

func NewPositionHandler(
	positionService *service.PositionService,
) *PositionHandler {
	return &PositionHandler{
		positionService: positionService,
	}
}

func (h *PositionHandler) Create(c *gin.Context) {
	var req dto.PositionCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	position, err := h.positionService.Create(c.Request.Context(), &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.NewPositionResponse(position)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (h *PositionHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	position, err := h.positionService.FindByID(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := dto.NewPositionResponse(position)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *PositionHandler) FindAll(c *gin.Context) {
	positions, err := h.positionService.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resList := dto.NewPositionResponses(positions)

	c.JSON(http.StatusOK, gin.H{"data": resList})
}

func (h *PositionHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	var req dto.PositionUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	position, err := h.positionService.Update(c.Request.Context(), id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.NewPositionResponse(position)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *PositionHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	position, err := h.positionService.Delete(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := dto.NewPositionResponse(position)
	c.JSON(http.StatusOK, gin.H{"data": res})
}
