package _type

type InvalidCredentialsError struct {
	BaseError
}

func NewInvalidCredentialsError(message string) BaseErrorInterface {
	return &InvalidCredentialsError{BaseError{Message: message}}
}
