package main

import (
	"ujian3/app"
	"ujian3/delivery"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterStructValidation(
			delivery.RegisterRequestValidation,
			delivery.RegisterRequest{},
		)
		v.RegisterValidation(
			"companyemail",
			delivery.CompanyEmail,
		)
	}

	application := app.InitializedApplication()
	application.Cleanup.Start()

	application.Server.Run(":8080")
}
