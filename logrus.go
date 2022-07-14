package abslog

import (
	"context"
	"fmt"
	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

func getLogrusLogger(logLevel LogLevel, encoder EncoderType) AbsLog {
	logrus.WithContext(context.Background())

	switch encoder {
	case JSONEncoder:
		logrus.SetFormatter(stackdriver.NewFormatter())
	case ConsoleEncoder:
	default:
		panic(fmt.Sprintf("Encoder type '%v' is not supported", encoder))
	}

	//logrus.SetFormatter(stackdriver.NewFormatter())
	logrus.SetLevel(getLogrusLevel(logLevel))
	logrus.SetReportCaller(true)

	return &absLog{
		debug:  logrus.Debug,
		debugf: logrus.Debugf,
		info:   logrus.Info,
		infof:  logrus.Infof,
		warn:   logrus.Warn,
		warnf:  logrus.Warnf,
		error:  logrus.Error,
		errorf: logrus.Errorf,
		fatal:  logrus.Fatal,
		fatalf: logrus.Fatalf,
		panic:  logrus.Panic,
		panicf: logrus.Panicf,
	}
}

func getLogrusLevel(logLevel LogLevel) logrus.Level {
	switch logLevel {
	case DebugLevel:
		return logrus.DebugLevel
	case InfoLevel:
		return logrus.InfoLevel
	case WarnLevel:
		return logrus.WarnLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	case FatalLevel:
		return logrus.FatalLevel
	case PanicLevel:
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}
