package abslog

import "fmt"

func init() {
	SetLoggerType(defaultLoggerType)
}

var Debug func(args ...interface{})
var Debugf func(format string, args ...interface{})
var Info func(args ...interface{})
var Infof func(format string, args ...interface{})
var Warn func(args ...interface{})
var Warnf func(format string, args ...interface{})
var Error func(args ...interface{})
var Errorf func(format string, args ...interface{})
var Panic func(args ...interface{})
var Panicf func(format string, args ...interface{})
var Fatal func(args ...interface{})
var Fatalf func(format string, args ...interface{})

func SetLoggerType(typ LoggerType) {
	var al AbsLog
	switch typ {
	case ZapLogger:
		al = getZapLogger(defaultLogLevel, defaultEncoderType)
	case LogrusLogger:
		al = getLogrusLogger(defaultLogLevel, defaultEncoderType)
	default:
		panic(fmt.Sprintf("Logger type '%v' is not supported", typ))
	}
	SetCustomLogger(al)
}

func SetCustomLogger(logger AbsLog) {
	Debug = logger.Debug
	Debugf = logger.Debugf
	Info = logger.Info
	Infof = logger.Infof
	Warn = logger.Warn
	Warnf = logger.Warnf
	Error = logger.Error
	Errorf = logger.Errorf
	Panic = logger.Panic
	Panicf = logger.Panicf
	Fatal = logger.Fatal
	Fatalf = logger.Fatalf
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
