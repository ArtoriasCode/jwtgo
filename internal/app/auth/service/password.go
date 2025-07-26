package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	customErr "jwtgo/internal/pkg/error/type"
	"jwtgo/pkg/logging"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
	hashingCost int
	globalSalt  string
	logger      *logging.Logger
}

func NewPasswordService(hashingCost int, globalSalt string, logger *logging.Logger) *PasswordService {
	return &PasswordService{
		hashingCost: hashingCost,
		globalSalt:  globalSalt,
		logger:      logger,
	}
}

func (s *PasswordService) GenerateSalt(length int) (string, customErr.BaseErrorIface) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		s.logger.Error("[PasswordService -> GenerateSalt -> Read]: ", err)
		return "", customErr.NewInternalServerError("Failed to generate salt")
	}

	return hex.EncodeToString(randomBytes), nil
}

func (s *PasswordService) HashPassword(password, localSalt string) (string, customErr.BaseErrorIface) {
	saltedPassword := localSalt + password + s.globalSalt

	hash := sha256.Sum256([]byte(saltedPassword))
	preHashedPassword := hex.EncodeToString(hash[:])

	bytes, err := bcrypt.GenerateFromPassword([]byte(preHashedPassword), s.hashingCost)
	if err != nil {
		s.logger.Error("[PasswordService -> HashPassword -> GenerateFromPassword]: ", err)
		return "", customErr.NewInternalServerError("Failed to generate hash")
	}

	return string(bytes), nil
}

func (s *PasswordService) VerifyPassword(plainPassword, hashedPassword, localSalt string) bool {
	saltedPassword := localSalt + plainPassword + s.globalSalt

	hash := sha256.Sum256([]byte(saltedPassword))
	preHashedPassword := hex.EncodeToString(hash[:])

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(preHashedPassword))
	if err != nil {
		s.logger.Error("[PasswordService -> VerifyPassword -> CompareHashAndPassword]: ", err)
	}

	return err == nil
}
