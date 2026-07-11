package errors

// Errors represents a slice of errors.
type Errors []error

// Add an existing error.
func (errs *Errors) Add(err error) {
	*errs = append(*errs, err)
}

// Add a new error with the given message.
func (errs *Errors) AddNew(msg string) {
	errs.Add(New(msg))
}

// Format the message with fmt then add the new error.
func (errs *Errors) AddFormat(format string, a ...any) {
	errs.Add(Format(format, a...))
}

// Collapse the errors into a single error object using the separator string.
// Returns nil if the slice is empty.
func (errs Errors) Collapse(sep string) error {
	if len(errs) == 0 {
		return nil
	}
	msg := ""

	for _, err := range errs {
		msg += err.Error() + sep
	}

	return New(msg[:len(msg)-len(sep)])
}
