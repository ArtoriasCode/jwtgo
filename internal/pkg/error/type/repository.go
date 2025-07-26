package _type

type AlreadyExistsError struct {
	BaseError
}

func NewAlreadyExistsError(message string) BaseErrorInterface {
	return &AlreadyExistsError{BaseError{Message: message}}
}

type NotFoundError struct {
	BaseError
}

func NewNotFoundError(message string) BaseErrorInterface {
	return &NotFoundError{BaseError{Message: message}}
}
