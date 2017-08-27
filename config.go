package logging

import "os"

var (
	DefaultFormatter Formatter = NewTimeFormatter("2006-01-02 15:04:05.000 -07:00", false)
	DefaultLogger    Logger    = NewColorTerm(os.Stdout, os.Stdout, os.Stderr, DefaultFormatter)
)
