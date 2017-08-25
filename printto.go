package logging

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var DefaultLogger Logger = NewPrintTo(os.Stderr, true, DefaultFormatter)

type PrintTo struct {
	writer    io.Writer
	debug     bool
	formatter Formatter
	mutex     sync.Mutex
}

func NewPrintTo(writer io.Writer, debug bool, formatter Formatter) *PrintTo {
	if formatter == nil {
		formatter = DefaultFormatter
	}
	return &PrintTo{writer: writer, debug: debug, formatter: formatter}
}

func (printTo *PrintTo) Printf(msg string, v ...interface{}) {
	printTo.mutex.Lock()
	defer printTo.mutex.Unlock()
	fmt.Fprintln(printTo.writer, printTo.formatter.Format(msg, v...))
}

func (printTo *PrintTo) Debugf(msg string, v ...interface{}) {
	if printTo.debug {
		printTo.mutex.Lock()
		defer printTo.mutex.Unlock()
		fmt.Fprintln(printTo.writer, printTo.formatter.Format(msg, v...))
	}
}

func (printTo *PrintTo) UnresolvedErrorf(err error, msg string, v ...interface{}) {
	printTo.mutex.Lock()
	defer printTo.mutex.Unlock()
	fmt.Fprintln(printTo.writer, printTo.formatter.FormatError(err, msg, v...))
}
