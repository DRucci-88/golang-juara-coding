package main

import "ujian2_rematch/app"

func main() {
	application := app.InitializedApplication()

	application.CleanUp.Start()

	application.Server.Run(":8080")
}
