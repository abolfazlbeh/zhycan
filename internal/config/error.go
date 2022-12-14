package config

// Imports needed list
import (
	"fmt"
)

// KeyNotExistErr Error
type KeyNotExistErr struct {
	Key      interface{}
	Category interface{}
	Err      error
}

// Error method - satisfying error interface
func (err *KeyNotExistErr) Error() string {
	return fmt.Sprintf("The key: '%v' does not exist in %v | %v", err.Key, err.Category, err.Err)
}

// NewKeyNotExistErr - return a new instance of KeyNotExistErr
func NewKeyNotExistErr(key, category string, err error) error {
	return &KeyNotExistErr{
		Key:      key,
		Category: category,
		Err:      err,
	}
}

// CategoryNotExistErr Error
type CategoryNotExistErr struct {
	Key interface{}
	Err error
}

// Error method - satisfying error interface
func (err *CategoryNotExistErr) Error() string {
	return fmt.Sprintf("The config file '%v' does not exist | %v", err.Key, err.Err)
}

// NewCategoryNotExistErr - return a new instance of CategoryNotExistErr
func NewCategoryNotExistErr(key string, err error) error {
	return &CategoryNotExistErr{
		Key: key,
		Err: err,
	}
}

// RemoteLoadErr Error
type RemoteLoadErr struct {
	Name interface{}
	Err  error
}

// Error method - satisfying error interface
func (err *RemoteLoadErr) Error() string {
	return fmt.Sprintf("Cannot load remote config: %v | %v", err.Name, err.Err)
}

// NewRemoteLoadErr - return a new instance of RemoteLoadErr
func NewRemoteLoadErr(name interface{}, err error) error {
	return &RemoteLoadErr{
		Name: name,
		Err:  err,
	}
}

// RemoteResponseErr Error
type RemoteResponseErr struct {
	Err error
}

// Error method - satisfying error interface
func (err *RemoteResponseErr) Error() string {
	return fmt.Sprintf("Cannot get response from remote | %v", err.Err)
}

// NewRemoteResponseErr - return a new instance of RemoteResponseErr
func NewRemoteResponseErr(err error) error {
	return &RemoteResponseErr{Err: err}
}
