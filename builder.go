package abslog

import "fmt"

type LogLevel int8

const (
	DebugLevel LogLevel = iota + 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

type LoggerType int8

const (
	ZapLogger LoggerType = iota + 1
	LogrusLogger
)
const defaultLoggerType = ZapLogger

type LoggerGen func(logLevel LogLevel) *AbsLog

// AbsLogBuilder is a builder for creating a new AbsLogger.
type AbsLogBuilder struct {
	logLevel   LogLevel
	loggerGen  LoggerGen
	loggerType LoggerType
}

// GetAbsLogBuilder returns a new AbsLogBuilder.
func GetAbsLogBuilder() *AbsLogBuilder {
	return &AbsLogBuilder{
		logLevel:   InfoLevel,
		loggerType: defaultLoggerType,
	}
}

// LogLevel sets the log level for the AbsLogger.
func (builder *AbsLogBuilder) LogLevel(level LogLevel) *AbsLogBuilder {
	builder.logLevel = level
	return builder
}

func (builder *AbsLogBuilder) LoggerGen(generator LoggerGen) *AbsLogBuilder {
	builder.loggerGen = generator
	return builder
}

func (builder *AbsLogBuilder) LoggerType(loggerType LoggerType) *AbsLogBuilder {
	builder.loggerType = loggerType
	return builder
}

// Build builds a new AbsLogger.
func (builder *AbsLogBuilder) Build() *AbsLog {
	if builder.loggerGen == nil {
		switch builder.loggerType {
		case ZapLogger:
			builder.loggerGen = GetZapLogger
		case LogrusLogger:
			builder.loggerGen = GetLogrusLogger
		default:
			panic(fmt.Sprintf("AbsLog type '%s' is not supported", builder.loggerType))
		}
	}
	return builder.loggerGen(builder.logLevel)
}
