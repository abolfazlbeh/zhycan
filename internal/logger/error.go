package logger

import "fmt"

// Error Object
type Error struct {
	Err error
}

// Error method - satisfying error interface
func (err *Error) Error() string {
	return fmt.Sprintf("Logger Error: %v", err.Err)
}

// NewError - return a new instance of Error
func NewError(err error) error {
	return &Error{Err: err}
}
