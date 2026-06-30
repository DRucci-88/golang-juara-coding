//go:build wireinject
// +build wireinject

package app

import (
	"praktikum/handler"
	"praktikum/repository"
	"praktikum/service"
	"praktikum/worker"

	"github.com/google/wire"
)

func InitializedApplication() *Application {

	wire.Build(
		// App
		NewApplication,
		NewDatabase,
		NewRouter,

		// Worker
		worker.NewTokenCleanupWorker,

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
