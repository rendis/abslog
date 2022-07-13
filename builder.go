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

type LoggerGen func(logLevel LogLevel) AbsLog

// absLogBuilder is a builder for creating a new AbsLogger.
type absLogBuilder struct {
	logLevel   LogLevel
	loggerGen  LoggerGen
	loggerType LoggerType
}

// GetAbsLogBuilder returns a new AbsLog builder.
func GetAbsLogBuilder() *absLogBuilder {
	return &absLogBuilder{
		logLevel:   InfoLevel,
		loggerType: defaultLoggerType,
	}
}

// LogLevel sets the log level for the AbsLog.
func (builder *absLogBuilder) LogLevel(level LogLevel) *absLogBuilder {
	builder.logLevel = level
	return builder
}

// LoggerGen sets the AbsLog generator function.
func (builder *absLogBuilder) LoggerGen(generator LoggerGen) *absLogBuilder {
	builder.loggerGen = generator
	return builder
}

// LoggerType sets the logger type for the AbsLog.
func (builder *absLogBuilder) LoggerType(loggerType LoggerType) *absLogBuilder {
	builder.loggerType = loggerType
	return builder
}

// Build builds a new AbsLog.
func (builder *absLogBuilder) Build() AbsLog {
	if builder.loggerGen == nil {
		switch builder.loggerType {
		case ZapLogger:
			builder.loggerGen = getZapLogger
		case LogrusLogger:
			builder.loggerGen = getLogrusLogger
		default:
			panic(fmt.Sprintf("AbsLog type '%s' is not supported", builder.loggerType))
		}
	}
	return builder.loggerGen(builder.logLevel)
}

// BuildAndSetAsGlobal builds a new AbsLogger and sets it as the global AbsLog.
func (builder *absLogBuilder) BuildAndSetAsGlobal() AbsLog {
	if builder.loggerGen == nil {
		switch builder.loggerType {
		case ZapLogger:
			builder.loggerGen = getZapLogger
		case LogrusLogger:
			builder.loggerGen = getLogrusLogger
		default:
			panic(fmt.Sprintf("AbsLog type '%s' is not supported", builder.loggerType))
		}
	}
	l := builder.loggerGen(builder.logLevel)
	SetGlobalLogger(l)
	return l
}
