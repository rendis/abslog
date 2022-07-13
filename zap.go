package abslog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

const logTimeFormat = "2006-01-02T15:04:05Z"

func getZapLogger(logLevel LogLevel) AbsLog {

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
	enc := zapcore.NewJSONEncoder(cfg)

	// Get ZapCore equivalent of log level
	zapLevel := getZapLevel(logLevel)

	// Stdout level enabler
	stdoutLevels := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapLevel && level < zap.ErrorLevel
	})

	// Stderr level enabler
	stderrLevels := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel && level >= zapLevel
	})

	// Write syncers
	stdoutSyncer := zapcore.Lock(os.Stdout)
	stderrSyncer := zapcore.Lock(os.Stderr)

	// Core multi-output
	core := zapcore.NewTee(
		zapcore.NewCore(
			enc,
			stdoutSyncer,
			stdoutLevels,
		),
		zapcore.NewCore(
			enc,
			stderrSyncer,
			stderrLevels,
		),
	)

	// Create logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	sugar := logger.Sugar()

	return &absLog{
		debug:  sugar.Debug,
		debugf: sugar.Debugf,
		info:   sugar.Info,
		infof:  sugar.Infof,
		warn:   sugar.Warn,
		warnf:  sugar.Warnf,
		error:  sugar.Error,
		errorf: sugar.Errorf,
		fatal:  sugar.Fatal,
		fatalf: sugar.Fatalf,
		panic:  sugar.Panic,
		panicf: sugar.Panicf,
	}
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(logTimeFormat))
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
