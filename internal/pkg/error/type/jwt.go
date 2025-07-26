package _type

type InvalidTokenError struct {
	BaseError
}

func NewInvalidTokenError(message string) BaseErrorInterface {
	return &InvalidTokenError{BaseError{Message: message}}
}

type ExpiredTokenError struct {
	BaseError
}

func NewExpiredTokenError(message string) BaseErrorInterface {
	return &ExpiredTokenError{BaseError{Message: message}}
}
