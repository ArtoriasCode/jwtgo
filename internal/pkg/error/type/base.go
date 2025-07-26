package _type

type BaseErrorInterface interface {
	error
}

type BaseError struct {
	Message string
}

func (e *BaseError) Error() string {
	return e.Message
}
