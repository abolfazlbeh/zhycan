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

// NotExistServiceNameErr Error
type NotExistServiceNameErr struct {
	serviceName string
}

// Error method - satisfying error interface
func (err *NotExistServiceNameErr) Error() string {
	return fmt.Sprintf("Instance with service name not exist: %v", err.serviceName)
}

// NewNotExistServiceNameErr - return a new instance of NotExistServiceNameErr
func NewNotExistServiceNameErr(serviceName string) error {
	return &NotExistServiceNameErr{serviceName: serviceName}
}

// MigrateErr Error
type MigrateErr struct {
	Err error
}

// Error method - satisfying error interface
func (err *MigrateErr) Error() string {
	return fmt.Sprintf("Migrating tables got error: %v", err.Err)
}

// NewMigrateErr - return a new instance of MigrateErr
func NewMigrateErr(err error) error {
	return &MigrateErr{Err: err}
}

// CreateMongoWrapperErr Error
type CreateMongoWrapperErr struct {
	Err error
}

// Error method - satisfying error interface
func (err *CreateMongoWrapperErr) Error() string {
	return fmt.Sprintf("Create a new mongo wrapper encounterred an error: %v", err.Err)
}

// NewCreateMongoWrapperErr - return a new instance of CreateMongoWrapperErr
func NewCreateMongoWrapperErr(err error) error {
	return &CreateMongoWrapperErr{Err: err}
}
