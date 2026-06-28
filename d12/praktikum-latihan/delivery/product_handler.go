package delivery

import (
	"errors"
	"net/http"
	"praktikum/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	SKU   string  `json:"sku" binding:"required,sku_format"` // Menggunakan custom validator 'sku_format'
	Name  string  `json:"name" binding:"required,min=3,max=100"`
	Price float64 `json:"price" binding:"required,gt=0"`
	Stock int     `json:"stock" binding:"required,min=0"`
}

// DTO Output Response (Untuk menyembunyikan kolom sensitif/tidak penting)
type ProductResponse struct {
	ID    uint    `json:"id"`
	SKU   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type ProductHandler struct {
	usecase domain.ProductUsecase
}

func NewProductHandler(r *gin.Engine, usecase domain.ProductUsecase) {
	handler := &ProductHandler{usecase: usecase}
	// Daftarkan endpoint ke Gin
	r.POST("/api/v1/products", handler.Create)
	r.GET("/api/v1/products/:id", handler.GetByID)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Validasi input gagal", "error": err.Error()})
		return
	}
	// Panggil usecase dengan context
	prod, err := h.usecase.Create(c.Request.Context(), req.SKU, req.Name, req.Price, req.Stock)
	if err != nil {
		statusCode, errMsg := mapError(err)
		c.JSON(statusCode, gin.H{"success": false, "message": errMsg})
		return
	}
	// Transformasi domain entity ke DTO Response
	response := ProductResponse{
		ID:    prod.ID,
		SKU:   prod.SKU,
		Name:  prod.Name,
		Price: prod.Price,
		Stock: prod.Stock,
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Produk berhasil dibuat", "data": response})
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	prod, err := h.usecase.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		statusCode, errMsg := mapError(err)
		c.JSON(statusCode, gin.H{"success": false, "message": errMsg})
		return
	}

	response := ProductResponse{
		ID:    prod.ID,
		SKU:   prod.SKU,
		Name:  prod.Name,
		Price: prod.Price,
		Stock: prod.Stock,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": response})
}

// Helper untuk standarisasi pemetaan error bisnis ke HTTP Code
func mapError(err error) (int, string) {
	if errors.Is(err, domain.ErrProductNotFound) {
		return http.StatusNotFound, err.Error()
	}
	if errors.Is(err, domain.ErrSKUDuplicate) {
		return http.StatusConflict, err.Error()
	}
	return http.StatusInternalServerError, "Terjadi kesalahan internal padaserver"
}
