package errors

import "fmt"

var separator = "\n  "

// Set the separator string used for chaining errors.
func SetChainSeparator(sep string) {
	separator = sep
}

// Chain the error behind a new message.
func Chain(err error, msg string) error {
	return Format("%s%s%s", msg, separator, err)
}

// Chain the error behind a new message formatted it with fmt.
func ChainFormat(err error, format string, a ...any) error {
	return Chain(err, fmt.Sprintf(format, a...))
}
