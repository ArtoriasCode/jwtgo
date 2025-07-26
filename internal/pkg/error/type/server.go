package _type

type InternalServerError struct {
	BaseError
}

func NewInternalServerError(message string) BaseErrorInterface {
	return &InternalServerError{BaseError{Message: message}}
}
