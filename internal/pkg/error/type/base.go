package _type

type BaseErrorIface interface {
	error
}

type BaseError struct {
	Message string
}

func (e *BaseError) Error() string {
	return e.Message
}
