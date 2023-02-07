package watcher

import "fmt"

// StartWatcherErr Error
type StartWatcherErr struct {
	Err error
}

// Error method - satisfying error interface
func (err *StartWatcherErr) Error() string {
	return fmt.Sprintf("Cannot start watcher instance | %v", err.Err)
}

// NewStartWatcherErr - return a new instance of StartWatcherErr
func NewStartWatcherErr(err error) error {
	return &StartWatcherErr{Err: err}
}
