package logging

type Tee []Logger

func (tee Tee) Printf(msg string, v ...interface{}) {
	for _, logger := range tee {
		logger.Printf(msg, v...)
	}
}

func (tee Tee) Debugf(msg string, v ...interface{}) {
	for _, logger := range tee {
		logger.Debugf(msg, v...)
	}
}

func (tee Tee) UnresolvedErrorf(err error, msg string, v ...interface{}) {
	for _, logger := range tee {
		logger.UnresolvedErrorf(err, msg, v...)
	}
}
