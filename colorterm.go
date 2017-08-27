package logging

import (
	"io"
	"sync"

	"github.com/fatih/color"
)

type ColorTerm struct {
	out          io.Writer
	debugOut     io.Writer
	errOut       io.Writer
	formatter    Formatter
	DefaultColor *color.Color
	DebugColor   *color.Color
	ErrorColor   *color.Color
	mutex        sync.Mutex
}

func NewColorTerm(out, debugOut, errOut io.Writer, formatter Formatter) *ColorTerm {
	return &ColorTerm{
		out:          out,
		debugOut:     debugOut,
		errOut:       errOut,
		formatter:    formatter,
		DefaultColor: color.New(color.FgWhite),
		DebugColor:   color.New(color.FgHiCyan),
		ErrorColor:   color.New(color.FgHiRed),
	}
}

func (t *ColorTerm) Printf(msg string, v ...interface{}) {
	if t.out == nil {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.DefaultColor.Fprintln(t.out, t.formatter.Format(msg, v...))
}

func (t *ColorTerm) Debugf(msg string, v ...interface{}) {
	if t.debugOut == nil {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.DebugColor.Fprintln(t.debugOut, t.formatter.Format(msg, v...))
}

func (t *ColorTerm) UnresolvedErrorf(err error, msg string, v ...interface{}) {
	if t.errOut == nil {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.ErrorColor.Fprintln(t.errOut, t.formatter.FormatError(err, msg, v...))
}
