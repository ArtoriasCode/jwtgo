package main

import (
	"jwtgo/internal/app/user"
)

func main() {
	userMicroservice := user.NewUserMicroservice()
	userMicroservice.Initialize()
	userMicroservice.Run()
}
