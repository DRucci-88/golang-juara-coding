package handler

import (
	"net/http"
	"strconv"
	"ujian2_rematch/dto"
	"ujian2_rematch/service"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	employeeService *service.EmployeeService
}

func NewEmployeeHandler(
	employeeService *service.EmployeeService,
) *EmployeeHandler {
	return &EmployeeHandler{
		employeeService: employeeService,
	}
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var req dto.EmployeeCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	employee, err := h.employeeService.Create(c.Request.Context(), &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := dto.NewEmployeeResponse(employee)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (h *EmployeeHandler) FindAll(c *gin.Context) {
	var filter dto.EmployeeFilterRequest

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	employees, err := h.employeeService.FindAll(c.Request.Context(), &filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := dto.NewEmployeeResponses((employees))
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (h *EmployeeHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi((c.Param("id")))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	employee, err := h.employeeService.FindByID(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := dto.NewEmployeeResponse(employee)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *EmployeeHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi((c.Param("id")))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ID param is not a number"})
		return
	}

	var req dto.EmployeeUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	employee, err := h.employeeService.Update(c.Request.Context(), id, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := dto.NewEmployeeResponse(employee)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}
