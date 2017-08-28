package logging

import "fmt"

// Printer can be used to connect to other logging systems.
// Note that the standard library's log.Logger implements Printer
type Printer interface {
	// Printf logs a message
	Printf(msg string, v ...interface{})
}

type ToPrinter struct {
	printer Printer
	debug   bool
}

func NewToPrinter(printer Printer, debug bool) *ToPrinter {
	return &ToPrinter{printer, debug}
}

func (wrapper *ToPrinter) Printf(msg string, v ...interface{}) {
	wrapper.printer.Printf(msg, v...)
}

func (wrapper *ToPrinter) Debug(msg string, v ...interface{}) {
	if !wrapper.debug {
		return
	}
	wrapper.printer.Printf(msg, v...)
}

func (wrapper *ToPrinter) UnresolvedErrorf(err error, msg string, v ...interface{}) {
	wrapper.printer.Printf("%s: %+v", fmt.Sprintf(msg, v...), err)
}
