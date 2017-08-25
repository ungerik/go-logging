package logging

type Logger interface {
	// Printf logs a message
	Printf(msg string, v ...interface{})

	// Debugf logs debugging information
	Debugf(msg string, v ...interface{})

	// UnresolvedErrorf logs errors that can't be resolved by the application
	UnresolvedErrorf(err error, msg string, v ...interface{})
}
