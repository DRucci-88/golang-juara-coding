//go:build wireinject
// +build wireinject

package app

import (
	"ujian2_rematch/handler"
	"ujian2_rematch/middleware"
	"ujian2_rematch/repository"
	"ujian2_rematch/service"

	"github.com/google/wire"
)

func InitializedApplication() *Application {
	wire.Build(
		// App
		NewApplication,
		NewDatabase,
		NewRouter,

		// Middleware
		middleware.NewMiddlewareManager,

		// Handler
		handler.NewAuthHandler,
		handler.NewDepartmentHandler,
		handler.NewPositionHandler,
		handler.NewEmployeeHandler,

		// Service
		service.NewAuthService,
		service.NewDepartmentService,
		service.NewPositionService,
		service.NewEmployeeService,

		// Repository Manager
		repository.NewRepositoryManager,
	)
	return nil
}
