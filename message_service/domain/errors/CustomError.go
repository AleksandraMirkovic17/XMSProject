package errors

type CustomError struct {
	message string
}

func NewCustomError(message string) *CustomError {
	return &CustomError{message}
}

func (err *CustomError) Error() string { return err.message }
