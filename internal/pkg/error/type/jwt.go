package _type

type InvalidTokenError struct {
	BaseError
}

func NewInvalidTokenError(message string) BaseErrorIface {
	return &InvalidTokenError{BaseError{Message: message}}
}

type ExpiredTokenError struct {
	BaseError
}

func NewExpiredTokenError(message string) BaseErrorIface {
	return &ExpiredTokenError{BaseError{Message: message}}
}
