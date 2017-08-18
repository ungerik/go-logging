package logging

type Tee []Logger

func (tee Tee) Printf(format string, args ...interface{}) {
	for _, logger := range tee {
		logger.Printf(format, args...)
	}
}

func (tee Tee) Debugf(format string, args ...interface{}) {
	for _, logger := range tee {
		logger.Debugf(format, args...)
	}
}
