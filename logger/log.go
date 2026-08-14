package logger

import "fmt"

func Log(msg string) {
	if logger == nil {
		return
	}

	logger.Println(msg)
}

func Logf(format string, a ...any) {
	Log(fmt.Sprintf(format, a...))
}

func LogError(err error) {
	Log(err.Error())
}
