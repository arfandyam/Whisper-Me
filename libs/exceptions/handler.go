package exceptions

import "fmt"

type CustomError struct {
	Status      int
	Description string
	Message     string
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("Error: %d: %s", ce.Status, ce.Message)
}

func NewCustomError(status int, description string, message string) *CustomError {
	return &CustomError{
		Status:  status,
		Description: description,
		Message: message,
	}
}
