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

type InvalidCredentialsError struct {
	message string
}

func NewInvalidCredentialsError(message string) error {
	return &InvalidCredentialsError{message: message}
}

func (e *InvalidCredentialsError) Error() string {
	return e.message
}

type UserNotFoundError struct {
	message string
}

func NewUserNotFoundError(message string) error {
	return &UserNotFoundError{message: message}
}

func (e *UserNotFoundError) Error() string {
	return e.message
}
