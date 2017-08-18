package logging

import (
	"log"
)

type StdWrapper struct {
	logger *log.Logger
	debug  bool
}

func NewStdWrapper(logger *log.Logger, debug bool) *StdWrapper {
	return &StdWrapper{logger, debug}
}

func (wrapper *StdWrapper) Printf(format string, args ...interface{}) {
	wrapper.logger.Printf(format, args...)
}

func (wrapper *StdWrapper) Debug(format string, args ...interface{}) {
	if wrapper.debug {
		wrapper.logger.Printf(format, args...)
	}
}
