package error

type InternalServerError struct {
	message string
}

func NewInternalServerError(message string) error {
	return &InternalServerError{message: message}
}

func (e *InternalServerError) Error() string {
	return e.message
}
