package error

type InvalidTokenError struct {
	message string
}

func NewInvalidTokenError(message string) error {
	return &InvalidTokenError{message: message}
}

func (e *InvalidTokenError) Error() string {
	return e.message
}

type ExpiredTokenError struct {
	message string
}

func NewExpiredTokenError(message string) error {
	return &ExpiredTokenError{message: message}
}

func (e *ExpiredTokenError) Error() string {
	return e.message
}
