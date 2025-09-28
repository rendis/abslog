package abslog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// logTimeFormat defines the time format used for log timestamps.
const logTimeFormat = "2006-01-02T15:04:05Z"

// getZapLogger creates and configures a Zap logger with the specified log level and encoder type.
// It sets up separate outputs for stdout (info and below) and stderr (error and above).
func getZapLogger(logLevel LogLevel, encoder EncoderType) AbsLog {

	// Encoder config
	cfg := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "severity",
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		TimeKey:       "timestamp",
		EncodeTime:    customTimeEncoder,
		CallerKey:     "caller",
		EncodeCaller:  zapcore.ShortCallerEncoder,
		StacktraceKey: "trace",
	}

	var enc zapcore.Encoder
	switch encoder {
	case ConsoleEncoder:
		enc = zapcore.NewConsoleEncoder(cfg)
	case JSONEncoder:
		enc = zapcore.NewJSONEncoder(cfg)
	default:
		panic(fmt.Sprintf("Encoder type '%v' is not supported", encoder))
	}

	// Get ZapCore equivalent of log level
	zapLevel := getZapLevel(logLevel)

	// Stdout level enabler: route info/warn/debug to stdout
	// Only logs at or above the specified level, but below error level
	stdoutLevels := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapLevel && level < zap.ErrorLevel
	})

	// Stderr level enabler: route error/fatal/panic to stderr
	// Only logs at error level and above, respecting the minimum log level
	stderrLevels := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel && level >= zapLevel
	})

	// Write syncers
	stdoutSyncer := zapcore.Lock(os.Stdout)
	stderrSyncer := zapcore.Lock(os.Stderr)

	// Core multi-output: combines stdout and stderr cores
	// This allows different log levels to be routed to appropriate outputs
	core := zapcore.NewTee(
		// Core for stdout (debug, info, warn)
		zapcore.NewCore(
			enc,
			stdoutSyncer,
			stdoutLevels,
		),
		// Core for stderr (error, fatal, panic)
		zapcore.NewCore(
			enc,
			stderrSyncer,
			stderrLevels,
		),
	)

	// Create logger with caller info and stack traces
	// AddCallerSkip(1) skips one frame to show the actual caller, not the wrapper
	// AddStacktrace(zap.ErrorLevel) adds stack traces for error and above
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	// Use sugar logger for easier variadic argument handling
	sugar := logger.Sugar()

	// Wrap in LoggerAdapter to implement the AbsLog interface
	return NewLoggerAdapter(sugar)
}

// customTimeEncoder formats time values using the predefined logTimeFormat.
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(logTimeFormat))
}

// getZapLevel converts an AbsLog LogLevel to the corresponding Zap log level.
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
