//go:build wireinject
// +build wireinject

package app

import (
	"ujian3/delivery"
	"ujian3/repository"
	"ujian3/usecase"

	"github.com/google/wire"
)

func InitializedApplication() *Application {

	wire.Build(
		// App
		NewApplication,
		NewDatabase,
		NewRouter,
		NewGroupMiddleware,

		// Handler
		delivery.NewAuthHandler,

		// Usecase
		usecase.NewAuthUsecase,

		// Repository
		repository.NewEmployeeRepository,
		repository.NewUserRepository,
		repository.NewBlackListedTokenRepository,
	)

	return nil
}
