package _type

type InternalServerError struct {
	BaseError
}

func NewInternalServerError(message string) BaseErrorIface {
	return &InternalServerError{BaseError{Message: message}}
}
