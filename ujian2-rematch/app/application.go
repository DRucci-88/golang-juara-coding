package app

import "github.com/gin-gonic/gin"

type Application struct {
	Server *gin.Engine
}

func NewApplication(
	server *gin.Engine,
) *Application {
	return &Application{
		Server: server,
	}
}
