package main

import (
	"jwtgo/internal/app/auth"
)

func main() {
	authMicroservice := auth.NewAuthMicroservice()
	authMicroservice.Initialize()
	authMicroservice.Run()
}
