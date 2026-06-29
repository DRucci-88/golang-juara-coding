package app

import "github.com/gin-gonic/gin"

type Application struct {
	Server *gin.Engine
	// Cleanup *worker.TokenCleanupWorker
}

func NewApplication(
	server *gin.Engine,
	// cleanup *worker.TokenCleanupWorker,
) *Application {
	return &Application{
		Server: server,
		// Cleanup: cleanup,
	}
}
