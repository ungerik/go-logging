package logging

import (
	"fmt"
	"io"
	"os"
)

var DefaultLogger Logger = NewPrintTo(os.Stderr, true, DefaultFormatter)

type PrintTo struct {
	writer    io.Writer
	debug     bool
	formatter Formatter
}

func NewPrintTo(writer io.Writer, debug bool, formatter Formatter) *PrintTo {
	if formatter == nil {
		formatter = DefaultFormatter
	}
	return &PrintTo{writer, debug, formatter}
}

func (printTo *PrintTo) Printf(format string, args ...interface{}) {
	fmt.Fprintln(printTo.writer, printTo.formatter.Format(format, args...))
}

func (printTo *PrintTo) Debugf(format string, args ...interface{}) {
	if printTo.debug {
		fmt.Fprintln(printTo.writer, printTo.formatter.Format(format, args...))
	}
}
