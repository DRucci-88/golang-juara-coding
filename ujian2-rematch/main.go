package main

import (
	"ujian2_rematch/app"
	"ujian2_rematch/dto"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterStructValidation(
			dto.LeaveRequestValidation,
			dto.LeaveCreateRequest{},
		)
	}

	application := app.InitializedApplication()

	application.CleanUp.Start()

	application.Server.Run(":8080")
}
