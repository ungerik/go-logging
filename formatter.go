package logging

import (
	"fmt"
	"time"
)

type Formatter interface {
	Format(msg string, v ...interface{}) string
	FormatDebug(msg string, v ...interface{}) string
	FormatError(err error, msg string, v ...interface{}) string
}

type TimeFormatter struct {
	TimeFormat string
	UTC        bool
}

func NewTimeFormatter(timeFormat string, utc bool) *TimeFormatter {
	return &TimeFormatter{TimeFormat: timeFormat, UTC: utc}
}

func (f *TimeFormatter) Format(msg string, v ...interface{}) string {
	t := time.Now()
	if f.UTC {
		t = t.UTC()
	}
	return t.Format(f.TimeFormat) + " " + fmt.Sprintf(msg, v...)
}

func (f *TimeFormatter) FormatDebug(msg string, v ...interface{}) string {
	return f.Format(msg, v...)
}

func (f *TimeFormatter) FormatError(err error, msg string, v ...interface{}) string {
	t := time.Now()
	if f.UTC {
		t = t.UTC()
	}
	return fmt.Sprintf("%s %s: %+v", t.Format(f.TimeFormat), fmt.Sprintf(msg, v...), err)
}
