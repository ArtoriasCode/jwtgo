package main

import (
	"jwtgo/internal/app"
)

func main() {
	ginApp := app.NewApplication()
	ginApp.Initialize()
	ginApp.Run()
}
