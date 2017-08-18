package logging

type Logger interface {
	Printf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}
