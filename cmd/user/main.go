package main

import (
	"jwtgo/internal/app/user"
)

func main() {
	userMicroservice := user.NewUserMicroService()
	userMicroservice.Initialize()
	userMicroservice.Run()
}
