package errors

import (
	"errors"
	"fmt"
)

// Create a new error with the given message.
func New(msg string) error {
	return errors.New(msg)
}

// Format the message with fmt to create a new error.
func Format(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}
