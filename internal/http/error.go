package http

import "fmt"

// NotImplementedErr Error
type NotImplementedErr struct {
}

// Error method - satisfying error interface
func (err *NotImplementedErr) Error() string {
	return fmt.Sprintf("Not Implemented Yet")
}

// NewNotImplementedErr - return a new instance of NotImplementedErr
func NewNotImplementedErr() error {
	return &NotImplementedErr{}
}

// CreateServerErr Error
type CreateServerErr struct {
	Err error
}

// Error method - satisfying error interface
func (err *CreateServerErr) Error() string {
	return fmt.Sprintf("Create a new server encounterred an error: %v", err.Err)
}

// NewCreateServerErr - return a new instance of CreateServerErr
func NewCreateServerErr(err error) error {
	return &CreateServerErr{Err: err}
}

// StartServerErr Error
type StartServerErr struct {
	Err           error
	ListenAddress string
}

// Error method - satisfying error interface
func (err *StartServerErr) Error() string {
	return fmt.Sprintf("Starting server on `%v` encounterred an error: %v", err.ListenAddress, err.Err)
}

// NewStartServerErr - return a new instance of StartServerErr
func NewStartServerErr(addr string, err error) error {
	return &StartServerErr{Err: err, ListenAddress: addr}
}

// ShutdownServerErr Error
type ShutdownServerErr struct {
	Err error
}

// Error method - satisfying error interface
func (err *ShutdownServerErr) Error() string {
	return fmt.Sprintf("Shutting down server encounterred an error: %v", err.Err)
}

// NewShutdownServerErr - return a new instance of StartServerErr
func NewShutdownServerErr(err error) error {
	return &ShutdownServerErr{Err: err}
}
