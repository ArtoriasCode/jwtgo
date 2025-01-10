package error

type AlreadyExistsError struct {
	message string
}

func NewAlreadyExistsError(message string) error {
	return &AlreadyExistsError{message: message}
}

func (e *AlreadyExistsError) Error() string {
	return e.message
}

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) error {
	return &NotFoundError{message: message}
}

func (e *NotFoundError) Error() string {
	return e.message
}
