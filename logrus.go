package abslog

import (
	"context"
	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

func GetLogrusLogger(logLevel LogLevel) *AbsLog {
	logrus.WithContext(context.Background())
	logrus.SetFormatter(stackdriver.NewFormatter())
	logrus.SetLevel(getLogrusLevel(logLevel))
	logrus.SetReportCaller(true)

	return &AbsLog{
		Debug:  logrus.Debug,
		Debugf: logrus.Debugf,
		Info:   logrus.Info,
		Infof:  logrus.Infof,
		Warn:   logrus.Warn,
		Warnf:  logrus.Warnf,
		Error:  logrus.Error,
		Errorf: logrus.Errorf,
		Fatal:  logrus.Fatal,
		Fatalf: logrus.Fatalf,
		Panic:  logrus.Panic,
		Panicf: logrus.Panicf,
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
