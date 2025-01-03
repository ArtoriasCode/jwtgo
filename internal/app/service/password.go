package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
	hashingCost int
	globalSalt  string
}

func NewPasswordService(hashingCost int, globalSalt string) *PasswordService {
	return &PasswordService{
		hashingCost: hashingCost,
		globalSalt:  globalSalt,
	}
}

func (pm *PasswordService) GenerateSalt(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(randomBytes), nil
}

func (pm *PasswordService) HashPassword(password, localSalt string) (string, error) {
	saltedPassword := localSalt + password + pm.globalSalt

	hash := sha256.Sum256([]byte(saltedPassword))
	preHashedPassword := hex.EncodeToString(hash[:])

	bytes, err := bcrypt.GenerateFromPassword([]byte(preHashedPassword), pm.hashingCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (pm *PasswordService) VerifyPassword(plainPassword, hashedPassword, localSalt string) bool {
	saltedPassword := localSalt + plainPassword + pm.globalSalt

	hash := sha256.Sum256([]byte(saltedPassword))
	preHashedPassword := hex.EncodeToString(hash[:])

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(preHashedPassword))
	return err == nil
}
