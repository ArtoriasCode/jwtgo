package _type

type AlreadyExistsError struct {
	BaseError
}

func NewAlreadyExistsError(message string) BaseErrorIface {
	return &AlreadyExistsError{BaseError{Message: message}}
}

type NotFoundError struct {
	BaseError
}

func NewNotFoundError(message string) BaseErrorIface {
	return &NotFoundError{BaseError{Message: message}}
}
