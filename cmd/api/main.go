package main

import (
	"jwtgo/internal/app/api"
)

func main() {
	apiGateway := auth.NewApiGateway()
	apiGateway.Initialize()
	apiGateway.Run()
}
