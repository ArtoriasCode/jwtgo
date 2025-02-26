package main

import (
	"jwtgo/internal/app/auth"
)

func main() {
	authMicroservice := auth.NewAuthMicroService()
	authMicroservice.Initialize()
	authMicroservice.Run()
}
