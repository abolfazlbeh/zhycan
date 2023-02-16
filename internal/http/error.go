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

// NotSupportedHttpMethodErr Error
type NotSupportedHttpMethodErr struct {
	method string
}

// Error method - satisfying error interface
func (err *NotSupportedHttpMethodErr) Error() string {
	return fmt.Sprintf("This method '%v' is not supported", err.method)
}

// NewNotSupportedHttpMethodErr - return a new instance of NotSupportedHttpMethodErr
func NewNotSupportedHttpMethodErr(method string) error {
	return &NotSupportedHttpMethodErr{method: method}
}

// AddRouteToNilServerErr Error
type AddRouteToNilServerErr struct {
	route string
}

// Error method - satisfying error interface
func (err *AddRouteToNilServerErr) Error() string {
	return fmt.Sprintf("There is no server to add route: %v", err.route)
}

// NewAddRouteToNilServerErr - return a new instance of AddRouteToNilServerErr
func NewAddRouteToNilServerErr(route string) error {
	return &AddRouteToNilServerErr{route: route}
}

// GetRouteByNameErr Error
type GetRouteByNameErr struct {
	name string
}

// Error method - satisfying error interface
func (err *GetRouteByNameErr) Error() string {
	return fmt.Sprintf("There is no route be the name: %v", err.name)
}

// NewGetRouteByNameErr - return a new instance of GetRouteByNameErr
func NewGetRouteByNameErr(name string) error {
	return &GetRouteByNameErr{name: name}
}

// FromNilServerErr Error
type FromNilServerErr struct {
}

// Error method - satisfying error interface
func (err *FromNilServerErr) Error() string {
	return fmt.Sprintf("There is no server to operate")
}

// NewFromNilServerErr - return a new instance of FromNilServerErr
func NewFromNilServerErr() error {
	return &FromNilServerErr{}
}

// FromMultipleServerErr Error
type FromMultipleServerErr struct {
}

// Error method - satisfying error interface
func (err *FromMultipleServerErr) Error() string {
	return fmt.Sprintf("There are more than one server to operate")
}

// NewFromMultipleServerErr - return a new instance of FromMultipleServerErr
func NewFromMultipleServerErr() error {
	return &FromMultipleServerErr{}
}
