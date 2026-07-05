package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"ujian2_rematch/dto"
	"ujian2_rematch/helper"
	"ujian2_rematch/service"

	"github.com/gin-gonic/gin"
)

type SalaryHandler struct {
	salaryService *service.SalaryService
}

func NewSalaryHandler(
	salaryService *service.SalaryService,
) *SalaryHandler {
	return &SalaryHandler{
		salaryService: salaryService,
	}
}

func (h *SalaryHandler) Calculate(c *gin.Context) {
	var req dto.SalaryCalculateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	val := strings.Split(req.Period, "-")
	year, _ := strconv.Atoi(val[0])
	month, _ := strconv.Atoi(val[1])
	period := helper.PayrollPeriod(year, time.Month(month))

	_, err := h.salaryService.Calculate(c.Request.Context(), period)
	log.Println(err)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Okay"})
}
