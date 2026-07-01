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

	department, err := h.positionService.Create(c.Request.Context(), &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res dto.PositionResponse
	res.FromModel(department)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (h *PositionHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	department, err := h.positionService.FindByID(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res dto.PositionResponse
	res.FromModel(department)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *PositionHandler) FindAll(c *gin.Context) {
	departments, err := h.positionService.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resList := make([]dto.PositionResponse, 0)
	for _, d := range departments {
		var res dto.PositionResponse
		res.FromModel(&d)
		resList = append(resList, res)
	}

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

	department, err := h.positionService.Update(c.Request.Context(),id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res dto.PositionResponse
	res.FromModel(department)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *PositionHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	department, err := h.positionService.Delete(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res dto.PositionResponse
	res.FromModel(department)
	c.JSON(http.StatusOK, gin.H{"data": res})
}
