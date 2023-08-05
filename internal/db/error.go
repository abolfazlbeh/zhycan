package db

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

// CreateSqlWrapperErr Error
type CreateSqlWrapperErr struct {
	Err error
}

// Error method - satisfying error interface
func (err *CreateSqlWrapperErr) Error() string {
	return fmt.Sprintf("Create a new sql wrapper encounterred an error: %v", err.Err)
}

// NewCreateSqlWrapperErr - return a new instance of CreateSqlWrapperErr
func NewCreateSqlWrapperErr(err error) error {
	return &CreateSqlWrapperErr{Err: err}
}

// NotSupportedDbTypeErr Error
type NotSupportedDbTypeErr struct {
	dbType string
}

// Error method - satisfying error interface
func (err *NotSupportedDbTypeErr) Error() string {
	return fmt.Sprintf("Not Supported Database Dialect: %v", err.dbType)
}

// NewNotSupportedDbTypeErr - return a new instance of NotSupportedDbTypeErr
func NewNotSupportedDbTypeErr(dbType string) error {
	return &NotSupportedDbTypeErr{dbType: dbType}
}
