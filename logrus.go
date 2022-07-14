package abslog

import (
	"context"
	"fmt"
	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

func getLogrusLogger(logLevel LogLevel, encoder EncoderType) AbsLog {
	logr := logrus.New()
	logr.WithContext(context.Background())

	switch encoder {
	case JSONEncoder:
		logr.SetFormatter(stackdriver.NewFormatter())
	case ConsoleEncoder:
	default:
		panic(fmt.Sprintf("Encoder type '%v' is not supported", encoder))
	}

	logr.SetLevel(getLogrusLevel(logLevel))
	logr.SetReportCaller(true)

	return logr
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
