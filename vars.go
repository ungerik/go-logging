package logging

type Vars struct {
	LoggerVar *Logger
	DebugVar  *bool
}

func (vars *Vars) Printf(msg string, v ...interface{}) {
	if vars.LoggerVar != nil && *vars.LoggerVar != nil {
		(*vars.LoggerVar).Printf(msg, v...)
	}
}

func (vars *Vars) Debugf(msg string, v ...interface{}) {
	if vars.LoggerVar != nil && *vars.LoggerVar != nil && vars.DebugVar != nil && *vars.DebugVar {
		(*vars.LoggerVar).Debugf(msg, v...)
	}
}

func (vars *Vars) UnresolvedErrorf(err error, msg string, v ...interface{}) {
	if vars.LoggerVar != nil && *vars.LoggerVar != nil {
		(*vars.LoggerVar).UnresolvedErrorf(err, msg, v...)
	}
}
