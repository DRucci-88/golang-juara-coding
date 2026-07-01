package main

import "ujian2_rematch/app"

func main() {
	application := app.InitializedApplication()

	application.Server.Run(":8080")
}
