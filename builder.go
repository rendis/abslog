package abslog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

type LogLevel int8

type LoggerGen func(logLevel LogLevel) AbsLog

const (
	DebugLevel LogLevel = iota + 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

// AbsLogBuilder is a builder for creating a new AbsLogger.
type AbsLogBuilder struct {
	logLevel  LogLevel
	loggerGen LoggerGen
}

// GetAbsLogBuilder returns a new AbsLogBuilder.
func GetAbsLogBuilder() *AbsLogBuilder {
	return &AbsLogBuilder{
		logLevel:  InfoLevel,
		loggerGen: defaultLoggerGen,
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

// Build builds a new AbsLogger.
func (builder *AbsLogBuilder) Build() AbsLog {
	if builder.loggerGen == nil {
		log.Fatalf("loggerGen is not set in AbsLogBuilder")
	}
	if builder.logLevel == 0 {
		log.Fatalf("logLevel is not set in AbsLogBuilder")
	}
	return builder.loggerGen(builder.logLevel)
}

func defaultLoggerGen(logLevel LogLevel) AbsLog {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// Get ZapCore equivalent of log level
	zapLevel := getZapLevel(logLevel)

	// Stdout level enabler
	stdoutLevels := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapLevel && level < zap.ErrorLevel
	})

	// Stderr level enabler
	stderrLevels := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapLevel
	})

	// Write syncers
	stdoutSyncer := zapcore.Lock(os.Stdout)
	stderrSyncer := zapcore.Lock(os.Stderr)

	// Core multi-output
	core := zapcore.NewTee(
		zapcore.NewCore(
			encoder,
			stdoutSyncer,
			stdoutLevels,
		),
		zapcore.NewCore(
			encoder,
			stderrSyncer,
			stderrLevels,
		),
	)

	// Create logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	return logger.Sugar()
}

func getZapLevel(logLevel LogLevel) zapcore.Level {
	switch logLevel {
	case DebugLevel:
		return zap.DebugLevel
	case InfoLevel:
		return zap.InfoLevel
	case WarnLevel:
		return zap.WarnLevel
	case ErrorLevel:
		return zap.ErrorLevel
	case PanicLevel:
		return zap.PanicLevel
	case FatalLevel:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
