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
	SetCustomLogger(l)
	return l
}

func (builder *absBuilder) build() AbsLog {
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

type alterAbsLog struct {
	al AbsLog
}

func (l *alterAbsLog) Debug(args ...interface{}) {
	l.al.Debug(args...)
}

func (l *alterAbsLog) Debugf(format string, args ...interface{}) {
	l.al.Debugf(format, args...)
}

func (l *alterAbsLog) Info(args ...interface{}) {
	l.al.Info(args...)
}

func (l *alterAbsLog) Infof(format string, args ...interface{}) {
	l.al.Infof(format, args...)
}

func (l *alterAbsLog) Warn(args ...interface{}) {
	l.al.Warn(args...)
}

func (l *alterAbsLog) Warnf(format string, args ...interface{}) {
	l.al.Warnf(format, args...)
}

func (l *alterAbsLog) Error(args ...interface{}) {
	l.al.Error(args...)
}

func (l *alterAbsLog) Errorf(format string, args ...interface{}) {
	l.al.Errorf(format, args...)
}

func (l *alterAbsLog) Panic(args ...interface{}) {
	l.al.Panic(args...)
}

func (l *alterAbsLog) Panicf(format string, args ...interface{}) {
	l.al.Panicf(format, args...)
}

func (l *alterAbsLog) Fatal(args ...interface{}) {
	l.al.Fatal(args...)
}

func (l *alterAbsLog) Fatalf(format string, args ...interface{}) {
	l.al.Fatalf(format, args...)
}
