package main

import (
	"latihan1/model"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

var categories = []model.Category{
	{ID: 1, Name: "Red", Code: "CAT-1"},
	{ID: 2, Name: "Green", Code: "CAT-2"},
	{ID: 3, Name: "Pink", Code: "CAT-2"},
}

var nextID = 4

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/categories", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"data": categories})
		})
		api.POST("/categories", func(ctx *gin.Context) {

			var category model.Category
			if err := ctx.ShouldBindJSON(&category); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if strings.Contains(category.Code, " ") {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Code cannot contains space"})
				return
			}
			category.ID = nextID
			nextID++
			categories = append(categories, category)

			ctx.JSON(http.StatusCreated, gin.H{
				"message": "Category Created",
				"data":    category,
			})
		})
		api.GET("/categories/:code", func(ctx *gin.Context) {
			code := ctx.Param("code")
			if idx := slices.IndexFunc(categories, func(c model.Category) bool {
				return c.Code == code
			}); idx != -1 {
				ctx.JSON(http.StatusOK, gin.H{"data": categories[idx]})
				return
			}
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		})
	}

	r.Run(":8080")
}
