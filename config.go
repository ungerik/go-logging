package logging

type Config struct {
	Logger   Logger
	DebugVar *bool
}

func (config *Config) Printf(format string, args ...interface{}) {
	if config.Logger != nil {
		config.Logger.Printf(format, args...)
	}
}

func (config *Config) Debugf(format string, args ...interface{}) {
	if config.Logger != nil && config.DebugVar != nil && *config.DebugVar {
		config.Logger.Debugf(format, args...)
	}
}
