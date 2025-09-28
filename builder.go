package abslog

import "fmt"

// LogLevel represents the severity level of log messages.
type LogLevel int8

// Log level constants defining the severity of log messages.
const (
	// DebugLevel is used for debug messages, typically only enabled during development.
	DebugLevel LogLevel = iota + 1
	// InfoLevel is used for general informational messages.
	InfoLevel
	// WarnLevel is used for warning messages that indicate potential issues.
	WarnLevel
	// ErrorLevel is used for error messages that indicate failures.
	ErrorLevel
	// PanicLevel is used for panic messages that cause the application to panic.
	PanicLevel
	// FatalLevel is used for fatal messages that cause the application to exit.
	FatalLevel
)

// EncoderType represents the format used for log output.
type EncoderType int8

// Encoder type constants defining the output format for log messages.
const (
	// ConsoleEncoder formats logs for human-readable console output.
	ConsoleEncoder EncoderType = iota + 1
	// JSONEncoder formats logs as JSON for structured logging.
	JSONEncoder
)

// LoggerType represents the underlying logging library to use.
type LoggerType int8

// Logger type constants defining which logging backend to use.
const (
	// ZapLogger uses the Uber Zap logging library as the backend.
	ZapLogger LoggerType = iota + 1
	// LogrusLogger uses the Sirupsen Logrus logging library as the backend.
	LogrusLogger
)

const defaultLogLevel = InfoLevel
const defaultLoggerType = ZapLogger
const defaultEncoderType = ConsoleEncoder

// LoggerGen is a function type that creates an AbsLog instance
// with the specified log level and encoder type.
type LoggerGen func(logLevel LogLevel, encoder EncoderType) AbsLog

// AbsLogBuilder is the interface that wraps the Builder methods to create a new AbsLog.
type AbsLogBuilder interface {
	LogLevel(level LogLevel) AbsLogBuilder
	LoggerGen(generator LoggerGen) AbsLogBuilder
	LoggerType(loggerType LoggerType) AbsLogBuilder
	EncoderType(encoderType EncoderType) AbsLogBuilder
	BuildAndSetAsGlobal() AbsLog
	Build() AbsLog
}

// absBuilder is a builder for creating a new AbsLogger.
type absBuilder struct {
	logLevel    LogLevel
	loggerGen   LoggerGen
	loggerType  LoggerType
	encoderType EncoderType
}

// GetAbsLogBuilder returns a new AbsLog builder.
func GetAbsLogBuilder() AbsLogBuilder {
	return &absBuilder{
		logLevel:    defaultLogLevel,
		loggerType:  defaultLoggerType,
		encoderType: defaultEncoderType,
	}
}

// LogLevel sets the log level for the AbsLog.
func (builder *absBuilder) LogLevel(level LogLevel) AbsLogBuilder {
	builder.logLevel = level
	return builder
}

// LoggerGen sets the AbsLog generator function.
func (builder *absBuilder) LoggerGen(generator LoggerGen) AbsLogBuilder {
	builder.loggerGen = generator
	return builder
}

// LoggerType sets the logger type for the AbsLog.
func (builder *absBuilder) LoggerType(loggerType LoggerType) AbsLogBuilder {
	builder.loggerType = loggerType
	return builder
}

// EncoderType sets the encoder type for the AbsLog.
func (builder *absBuilder) EncoderType(encoderType EncoderType) AbsLogBuilder {
	builder.encoderType = encoderType
	return builder
}

// Build builds a new AbsLog.
func (builder *absBuilder) Build() AbsLog {
	return builder.build()
}

// BuildAndSetAsGlobal builds a new AbsLogger and sets it as the global AbsLog.
func (builder *absBuilder) BuildAndSetAsGlobal() AbsLog {
	l := builder.build()
	SetLogger(l)
	return l
}

// build creates the final AbsLog instance using the configured settings.
// It validates the encoder type and sets up the appropriate logger generator if needed.
func (builder *absBuilder) build() AbsLog {
	// Validate encoder type
	if builder.encoderType != ConsoleEncoder && builder.encoderType != JSONEncoder {
		panic(fmt.Sprintf("Invalid encoder type: %d", builder.encoderType))
	}

	// Set default logger generator if not provided
	if builder.loggerGen == nil {
		switch builder.loggerType {
		case ZapLogger:
			builder.loggerGen = getZapLogger
		case LogrusLogger:
			builder.loggerGen = getLogrusLogger
		default:
			panic(fmt.Sprintf("AbsLog type '%d' is not supported", int(builder.loggerType)))
		}
	}

	// Create and return the logger instance
	return builder.loggerGen(builder.logLevel, builder.encoderType)
}
