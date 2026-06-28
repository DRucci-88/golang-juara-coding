package main

import "praktikum/app"

func main() {
	application := app.InitializedApplication()

	application.Cleanup.Start()
	application.Server.Run(":8080")

}
