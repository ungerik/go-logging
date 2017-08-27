package logging

import (
	"fmt"
	"io"
	"sync"
)

type PrintTo struct {
	out       io.Writer
	debugOut  io.Writer
	errOut    io.Writer
	formatter Formatter
	mutex     sync.Mutex
}

func NewPrintTo(out, debugOut, errOut io.Writer, formatter Formatter) *PrintTo {
	return &PrintTo{
		out:       out,
		debugOut:  debugOut,
		errOut:    errOut,
		formatter: formatter,
	}
}

func (p *PrintTo) Printf(msg string, v ...interface{}) {
	if p.out == nil {
		return
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	fmt.Fprintln(p.out, p.formatter.Format(msg, v...))
}

func (p *PrintTo) Debugf(msg string, v ...interface{}) {
	if p.debugOut == nil {
		return
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	fmt.Fprintln(p.debugOut, p.formatter.Format(msg, v...))
}

func (p *PrintTo) UnresolvedErrorf(err error, msg string, v ...interface{}) {
	if p.errOut == nil {
		return
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	fmt.Fprintln(p.errOut, p.formatter.FormatError(err, msg, v...))
}
