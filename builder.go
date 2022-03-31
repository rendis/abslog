package abslog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

type AbsLogBuilder struct {
	logLevel Level
}

var logGen = func(logLevel Level) Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	var core zapcore.Core
	switch logLevel {
	case DebugLevel:
		core = zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	case WarnLevel:
		core = zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), zap.WarnLevel)
	case ErrorLevel:
		core = zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), zap.ErrorLevel)
	case PanicLevel:
		core = zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), zap.PanicLevel)
	case FatalLevel:
		core = zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), zap.FatalLevel)
	default:
		core = zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), zap.InfoLevel)
	}
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	return logger.Sugar()
}

func SetLoggerGenerator(generator func(Level) Logger) {
	logGen = generator
}

func GetAbsLogBuilder() *AbsLogBuilder {
	return &AbsLogBuilder{InfoLevel}
}

func (builder *AbsLogBuilder) LogLevel(logLevel Level) *AbsLogBuilder {
	builder.logLevel = logLevel
	return builder
}

func (builder *AbsLogBuilder) Build() *AbsLog {
	return &AbsLog{
		logger: logGen(builder.logLevel),
	}
}
