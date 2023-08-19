package grpc

import "fmt"

// GrpcServerStartError struct
type GrpcServerStartError struct {
	Err error
}

// Error method - satisfying error interface
func (err *GrpcServerStartError) Error() string {
	return fmt.Sprintf("gRPC Start Server Error | %v", err.Err)
}

// NewGrpcServerStartError - return a new instance of GrpcServerStartError
func NewGrpcServerStartError(err error) error {
	return &GrpcServerStartError{Err: err}
}

// GrpcDialError struct
type GrpcDialError struct {
	addr string
	Err  error
}

// Error method - satisfying error interface
func (err *GrpcDialError) Error() string {
	return fmt.Sprintf("gRPC dial to %v Error | %v", err.addr, err.Err)
}

// NewGrpcDialError - return a new instance of GrpcDialError
func NewGrpcDialError(addr string, err error) error {
	return &GrpcDialError{Err: err, addr: addr}
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

// GrpcServerNotExistError struct
type GrpcServerNotExistError struct {
	name string
}

// Error method - satisfying error interface
func (err *GrpcServerNotExistError) Error() string {
	return fmt.Sprintf("gRPC server with specified name `%v` not exist", err.name)
}

// NewGrpcServerNotExistError - return a new instance of GrpcServerNotExistError
func NewGrpcServerNotExistError(name string) error {
	return &GrpcServerNotExistError{name: name}
}

// NilServiceRegistryError struct
type NilServiceRegistryError struct {
}

// Error method - satisfying error interface
func (err *NilServiceRegistryError) Error() string {
	return fmt.Sprintf("Register a Nil service to gRPC server")
}

// NewNilServiceRegistryError - return a new instance of NilServiceRegistryError
func NewNilServiceRegistryError() error {
	return &NilServiceRegistryError{}
}
