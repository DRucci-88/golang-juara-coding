package handler

import (
	"errors"
	"net/http"
	"strconv"

	"materi-middleware-gorm/models"
	"materi-middleware-gorm/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ConcertHandler struct {
	service service.ConcertService
}

func NewConcertHandler(service service.ConcertService) *ConcertHandler {
	return &ConcertHandler{service: service}
}

// CreateConcert handles POST /api/v1/concerts
func (h *ConcertHandler) CreateConcert(c *gin.Context) {
	var concert models.Concert
	if err := c.ShouldBindJSON(&concert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateConcert(&concert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create concert: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, concert)
}

// GetConcerts handles GET /api/v1/concerts
func (h *ConcertHandler) GetConcerts(c *gin.Context) {
	concerts, err := h.service.GetAllConcerts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve concerts: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, concerts)
}

// GetConcertByID handles GET /api/v1/concerts/:id
func (h *ConcertHandler) GetConcertByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	concert, err := h.service.GetConcertByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Concert not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, concert)
}

// UpdateConcert handles PUT /api/v1/concerts/:id
func (h *ConcertHandler) UpdateConcert(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	concert, err := h.service.GetConcertByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Concert not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input models.Concert
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	concert.Title = input.Title
	concert.Description = input.Description
	concert.Date = input.Date
	concert.Venue = input.Venue
	concert.Status = input.Status

	if err := h.service.UpdateConcert(&concert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update concert: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, concert)
}

// DeleteConcert handles DELETE /api/v1/concerts/:id
func (h *ConcertHandler) DeleteConcert(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.service.DeleteConcert(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Concert not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Concert deleted successfully"})
}
