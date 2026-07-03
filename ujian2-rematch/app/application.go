package app

import (
	"ujian2_rematch/worker"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Server  *gin.Engine
	CleanUp *worker.TokenCleanupWorker
}

func NewApplication(
	server *gin.Engine,
	cleanup *worker.TokenCleanupWorker,
) *Application {
	return &Application{
		Server:  server,
		CleanUp: cleanup,
	}
}
