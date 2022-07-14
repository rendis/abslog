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

type EncoderType int8

const (
	ConsoleEncoder EncoderType = iota + 1
	JSONEncoder
)

type LoggerType int8

const (
	ZapLogger LoggerType = iota + 1
	LogrusLogger
)

const defaultLogLevel = InfoLevel
const defaultLoggerType = ZapLogger
const defaultEncoderType = ConsoleEncoder

type LoggerGen func(logLevel LogLevel, encoder EncoderType) AbsLog

// absLogBuilder is a builder for creating a new AbsLogger.
type absLogBuilder struct {
	logLevel    LogLevel
	loggerGen   LoggerGen
	loggerType  LoggerType
	encoderType EncoderType
}

// GetAbsLogBuilder returns a new AbsLog builder.
func GetAbsLogBuilder() *absLogBuilder {
	return &absLogBuilder{
		logLevel:    defaultLogLevel,
		loggerType:  defaultLoggerType,
		encoderType: defaultEncoderType,
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

func (builder *absLogBuilder) EncoderType(encoderType EncoderType) *absLogBuilder {
	builder.encoderType = encoderType
	return builder
}

// Build builds a new AbsLog.
func (builder *absLogBuilder) Build() AbsLog {
	return builder.build()
}

// BuildAndSetAsGlobal builds a new AbsLogger and sets it as the global AbsLog.
func (builder *absLogBuilder) BuildAndSetAsGlobal() AbsLog {
	l := builder.build()
	SetGlobalLogger(l)
	return l
}

func (builder *absLogBuilder) build() AbsLog {
	if builder.encoderType != ConsoleEncoder && builder.encoderType != JSONEncoder {
		panic(fmt.Sprintf("Invalid encoder type: %d", builder.encoderType))
	}

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

	return builder.loggerGen(builder.logLevel, builder.encoderType)
}
