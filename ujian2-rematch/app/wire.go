//go:build wireinject
// +build wireinject

package app

import (
	"ujian2_rematch/handler"
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

		// Handler
		handler.NewDepartmentHandler,
		handler.NewPositionHandler,

		// Service
		service.NewDepartmentService,
		service.NewPositionService,

		// Repository Manager
		repository.NewRepositoryManager,
	)
	return nil
}
