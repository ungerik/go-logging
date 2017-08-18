package logging

import (
	"fmt"
	"time"
)

type Formatter interface {
	Format(format string, args ...interface{}) string
}

var DefaultFormatter Formatter = NewTimeFormatter("2006-01-02 15:04:05.999 -0700 MST", false)

type TimeFormatter struct {
	TimeFormat string
	UTC        bool
}

func NewTimeFormatter(timeFormat string, utc bool) *TimeFormatter {
	return &TimeFormatter{TimeFormat: timeFormat, UTC: utc}
}

func (f *TimeFormatter) Format(format string, args ...interface{}) string {
	t := time.Now()
	if f.UTC {
		t = t.UTC()
	}
	return t.Format(f.TimeFormat) + " " + fmt.Sprintf(format, args...)
}
