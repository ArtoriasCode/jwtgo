package service

import customErr "jwtgo/internal/pkg/error/type"

type PasswordServiceIface interface {
	GenerateSalt(length int) (string, customErr.BaseErrorIface)
	HashPassword(password, localSalt string) (string, customErr.BaseErrorIface)
	VerifyPassword(plainPassword, hashedPassword, localSalt string) bool
}
