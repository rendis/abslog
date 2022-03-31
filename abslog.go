package abslog

// Logger generic interface for logging
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

// AbsLog logger abstraction
type AbsLog struct {
	logger Logger
}

func (l *AbsLog) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *AbsLog) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *AbsLog) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *AbsLog) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *AbsLog) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *AbsLog) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *AbsLog) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *AbsLog) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *AbsLog) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *AbsLog) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}
