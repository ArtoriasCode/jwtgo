package error

type InvalidCredentialsError struct {
	message string
}

func NewInvalidCredentialsError(message string) error {
	return &InvalidCredentialsError{message: message}
}

func (e *InvalidCredentialsError) Error() string {
	return e.message
}
