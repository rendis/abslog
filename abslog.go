package abslog

var al AbsLog = getZapLogger(InfoLevel)

func SetGlobalLogger(logger AbsLog) {
	al = logger
}

type AbsLog interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type absLog struct {
	debug  func(args ...interface{})
	debugf func(format string, args ...interface{})

	info  func(args ...interface{})
	infof func(format string, args ...interface{})

	warn  func(args ...interface{})
	warnf func(format string, args ...interface{})

	error  func(args ...interface{})
	errorf func(format string, args ...interface{})

	panic  func(args ...interface{})
	panicf func(format string, args ...interface{})

	fatal  func(args ...interface{})
	fatalf func(format string, args ...interface{})
}

func (l *absLog) Debug(args ...interface{}) {
	l.debug(args...)
}

func (l *absLog) Debugf(format string, args ...interface{}) {
	l.debugf(format, args...)
}

func (l *absLog) Info(args ...interface{}) {
	l.info(args...)
}

func (l *absLog) Infof(format string, args ...interface{}) {
	l.infof(format, args...)
}

func (l *absLog) Warn(args ...interface{}) {
	l.warn(args...)
}

func (l *absLog) Warnf(format string, args ...interface{}) {
	l.warnf(format, args...)
}

func (l *absLog) Error(args ...interface{}) {
	l.error(args...)
}

func (l *absLog) Errorf(format string, args ...interface{}) {
	l.errorf(format, args...)
}

func (l *absLog) Panic(args ...interface{}) {
	l.panic(args...)
}

func (l *absLog) Panicf(format string, args ...interface{}) {
	l.panicf(format, args...)
}

func (l *absLog) Fatal(args ...interface{}) {
	l.fatal(args...)
}

func (l *absLog) Fatalf(format string, args ...interface{}) {
	l.fatalf(format, args...)
}

func Debug(args ...interface{}) {
	al.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	al.Debugf(format, args...)
}

func Info(args ...interface{}) {
	al.Info(args...)
}

func Infof(format string, args ...interface{}) {
	al.Infof(format, args...)
}

func Warn(args ...interface{}) {
	al.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	al.Warnf(format, args...)
}

func Error(args ...interface{}) {
	al.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	al.Errorf(format, args...)
}

func Panic(args ...interface{}) {
	al.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	al.Panicf(format, args...)
}

func Fatal(args ...interface{}) {
	al.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	al.Fatalf(format, args...)
}
