package errors

import "fmt"

var separator = "\n  "

// Set the separator string used for chaining errors.
func SetChainSeparator(sep string) {
	separator = sep
}

// Create a new error with the given message, then chain behind the first error.
func Chain(err error, msg string) error {
	return Format("%s%s%s", msg, separator, err)
}

// Format the message with fmt to create a new error, then chain behind the first error.
func ChainFormat(err error, format string, a ...any) error {
	return Chain(err, fmt.Sprintf(format, a...))
}
