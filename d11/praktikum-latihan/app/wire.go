//go:build wireinject
// +build wireinject

package app

import (
	"praktikum/handler"
	"praktikum/repository"
	"praktikum/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializedServer() *gin.Engine {
	wire.Build(
		// App
		NewRouter,
		NewDatabase,

		// Handler
		handler.NewOrderHandler,

		// Service
		service.NewOrderService,

		// Repository
		repository.NewOrderRepository,
		repository.NewProductRepository,
		repository.NewUserReposutory,
		repository.NewOrderItemRepository,
	)
	return nil
}
