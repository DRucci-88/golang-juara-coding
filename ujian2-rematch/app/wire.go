//go:build wireinject
// +build wireinject

package app

import (
	"ujian2_rematch/handler"
	"ujian2_rematch/middleware"
	"ujian2_rematch/repository"
	"ujian2_rematch/service"
	"ujian2_rematch/worker"

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

		// Worker
		worker.NewTokenCleanupWorker,

		// Handler
		handler.NewAuthHandler,
		handler.NewDepartmentHandler,
		handler.NewPositionHandler,
		handler.NewEmployeeHandler,
		handler.NewAttendanceHandler,
		handler.NewLeaveHandler,

		// Service
		service.NewAuthService,
		service.NewDepartmentService,
		service.NewPositionService,
		service.NewEmployeeService,
		service.NewAttendanceService,
		service.NewLeaveService,

		// Repository Manager
		repository.NewRepositoryManager,
	)
	return nil
}
