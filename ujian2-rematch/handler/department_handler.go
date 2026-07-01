package handler

import (
	"net/http"
	"strconv"
	"ujian2_rematch/dto"
	"ujian2_rematch/service"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	departmentService *service.DepartmentService
}

func NewDepartmentHandler(
	departmentService *service.DepartmentService,
) *DepartmentHandler {
	return &DepartmentHandler{
		departmentService: departmentService,
	}
}

func (h *DepartmentHandler) Create(c *gin.Context) {
	var req dto.DepartmentCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	department, err := h.departmentService.Create(c.Request.Context(), &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.NewDepartmentResponse(department)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (h *DepartmentHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	department, err := h.departmentService.FindByID(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := dto.NewDepartmentResponse(department)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *DepartmentHandler) FindAll(c *gin.Context) {
	departments, err := h.departmentService.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resList := dto.NewDepartmentResponses(departments)

	c.JSON(http.StatusOK, gin.H{"data": resList})
}

func (h *DepartmentHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	var req dto.DepartmentUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	department, err := h.departmentService.Update(c.Request.Context(), id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.NewDepartmentResponse(department)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *DepartmentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	department, err := h.departmentService.Delete(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := dto.NewDepartmentResponse(department)
	c.JSON(http.StatusOK, gin.H{"data": res})
}
