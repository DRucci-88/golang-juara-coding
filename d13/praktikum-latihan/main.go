package main

import (
	"praktikum/app"
)

func main() {
	r := app.InitializedServer()

	r.Run(":8080")
}
