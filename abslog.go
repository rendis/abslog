package abslog

type AbsLog struct {
	Debug  func(args ...interface{})
	Debugf func(format string, args ...interface{})

	Info  func(args ...interface{})
	Infof func(format string, args ...interface{})

	Warn  func(args ...interface{})
	Warnf func(format string, args ...interface{})

	Error  func(args ...interface{})
	Errorf func(format string, args ...interface{})

	Panic  func(args ...interface{})
	Panicf func(format string, args ...interface{})

	Fatal  func(args ...interface{})
	Fatalf func(format string, args ...interface{})
}
