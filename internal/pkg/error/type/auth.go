package _type

type InvalidCredentialsError struct {
	BaseError
}

func NewInvalidCredentialsError(message string) BaseErrorIface {
	return &InvalidCredentialsError{BaseError{Message: message}}
}
