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

		// Middleware
		NewGroupMiddleware,

		// Handler
		handler.NewAuthHandler,
		handler.NewOrderHandler,
		handler.NewAnaliticsHandler,

		// Service
		service.NewAuthService,
		service.NewOrderService,
		service.NewOrderItemService,

		// Repository
		repository.NewOrderRepository,
		repository.NewProductRepository,
		repository.NewUserRepository,
		repository.NewOrderItemRepository,
		repository.NewBlackListedTokenRepository,
	)
	return nil
}
