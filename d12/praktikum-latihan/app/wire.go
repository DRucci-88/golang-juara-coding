//go:build wireinject
// +build wireinject

package app

import (
	"praktikum/delivery"
	"praktikum/repository"
	"praktikum/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeServer() *gin.Engine {
	wire.Build(
		// App
		NewRouter,
		NewDatabase,

		// Delivery
		delivery.NewProductHandler,

		// Repository
		repository.NewProductRepository,

		// Usecase
		usecase.NewProductUsecase,
	)
	return nil
}
