package logger

import "log"

func New() *log.Logger {
	return log.Default()
}
