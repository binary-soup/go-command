package logger

import "fmt"

// Log a message. Does nothing if the logger isn't set.
func Log(msg string) {
	if logger == nil {
		return
	}

	logger.Println(msg)
}

// Log a message formatted with fmt. Does nothing if the logger isn't set.
func Logf(format string, a ...any) {
	Log(fmt.Sprintf(format, a...))
}

// Log an error message. Does nothing if the logger isn't set.
func LogError(err error) {
	Log(err.Error())
}
