package logging

import (
	"fmt"
	"log"
)

type StdWrapper struct {
	stdLogger *log.Logger
	debug     bool
}

func NewStdWrapper(stdLogger *log.Logger, debug bool) *StdWrapper {
	return &StdWrapper{stdLogger, debug}
}

func (wrapper *StdWrapper) Printf(msg string, v ...interface{}) {
	wrapper.stdLogger.Printf(msg, v...)
}

func (wrapper *StdWrapper) Debug(msg string, v ...interface{}) {
	if !wrapper.debug {
		return
	}
	wrapper.stdLogger.Printf(msg, v...)
}

func (wrapper *StdWrapper) UnresolvedErrorf(err error, msg string, v ...interface{}) {
	wrapper.stdLogger.Printf("%s: %+v", fmt.Sprintf(msg, v...), err)
}
