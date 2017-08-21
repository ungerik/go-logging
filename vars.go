package logging

type Vars struct {
	LoggerVar *Logger
	DebugVar  *bool
}

func (v *Vars) Printf(format string, args ...interface{}) {
	if v.LoggerVar != nil && *v.LoggerVar != nil {
		(*v.LoggerVar).Printf(format, args...)
	}
}

func (v *Vars) Debugf(format string, args ...interface{}) {
	if v.LoggerVar != nil && *v.LoggerVar != nil && v.DebugVar != nil && *v.DebugVar {
		(*v.LoggerVar).Debugf(format, args...)
	}
}
