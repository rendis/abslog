// Package abslog provides an abstraction layer for logging libraries,
// allowing seamless switching between different logging backends (Zap, Logrus)
// while maintaining a consistent API.
package abslog

import (
	"context"
	"fmt"
	"strings"
)

// Default values for context configuration
const (
	// defaultContextKey is the default key used to store context values
	defaultContextKey = "abslog"
	// defaultContextSeparator is the default separator between context and log messages
	defaultContextSeparator = " -> "
)

// contextKey is the current key used to store context values in context.Context
var contextKey = defaultContextKey

// contextSeparator is the current string used to separate context values from log messages
var contextSeparator = defaultContextSeparator

// contextFormatTemplate is the template used to format context values
const contextFormatTemplate = "[%s]%s"

// AbsLog defines the interface for abstracted logging functionality.
// It provides methods for logging at different levels with optional formatting.
type AbsLog interface {
	Debug(args ...any)
	Debugf(format string, args ...any)

	Info(args ...any)
	Infof(format string, args ...any)

	Warn(args ...any)
	Warnf(format string, args ...any)

	Error(args ...any)
	Errorf(format string, args ...any)

	Fatal(args ...any)
	Fatalf(format string, args ...any)

	Panic(args ...any)
	Panicf(format string, args ...any)
}

func init() {
	fmt.Println("init abslog with default logger type (zap)")
	SetLoggerType(defaultLoggerType)
}

// SetCtxKey sets the key used to retrieve context values from context.Context.
// This allows customization of how context values are stored and retrieved.
// If key is empty or only whitespace, the default key "abslog" will be used.
func SetCtxKey(key string) {
	key = strings.TrimSpace(key)
	if key == "" {
		contextKey = defaultContextKey
	} else {
		contextKey = key
	}
}

// GetCtxKey returns the current key used to retrieve context values from context.Context.
func GetCtxKey() string {
	return contextKey
}

// ResetCtxKey resets the context key to its default value.
func ResetCtxKey() {
	contextKey = defaultContextKey
}

// GetCtxSeparator returns the current separator used between context values and log messages.
func GetCtxSeparator() string {
	return contextSeparator
}

// ResetCtxSeparator resets the context separator to its default value.
func ResetCtxSeparator() {
	contextSeparator = defaultContextSeparator
}

// SetCtxSeparator sets the string used to separate context values from log messages.
// If separator is empty or only whitespace, the default separator " -> " will be used.
func SetCtxSeparator(separator string) {
	trimmed := strings.TrimSpace(separator)
	if trimmed == "" {
		contextSeparator = defaultContextSeparator
	} else {
		contextSeparator = separator
	}
}

// Debug logs a message at level Debug on the standard logger.
var Debug func(args ...any)
var DebugCtx func(ctx context.Context, args ...any)
var Debugf func(format string, args ...any)
var DebugCtxf func(ctx context.Context, format string, args ...any)

// Info logs a message at level Info on the standard logger.
var Info func(args ...any)
var InfoCtx func(ctx context.Context, args ...any)
var Infof func(format string, args ...any)
var InfoCtxf func(ctx context.Context, format string, args ...any)

// Warn logs a message at level Warn on the standard logger.
var Warn func(args ...any)
var WarnCtx func(ctx context.Context, args ...any)
var Warnf func(format string, args ...any)
var WarnCtxf func(ctx context.Context, format string, args ...any)

// Error logs a message at level Error on the standard logger.
var Error func(args ...any)
var ErrorCtx func(ctx context.Context, args ...any)
var Errorf func(format string, args ...any)
var ErrorCtxf func(ctx context.Context, format string, args ...any)

// Fatal logs a message at level Fatal on the standard logger.
var Fatal func(args ...any)
var FatalCtx func(ctx context.Context, args ...any)
var Fatalf func(format string, args ...any)
var FatalCtxf func(ctx context.Context, format string, args ...any)

// Panic logs a message at level Panic on the standard logger.
var Panic func(args ...any)
var PanicCtx func(ctx context.Context, args ...any)
var Panicf func(format string, args ...any)
var PanicCtxf func(ctx context.Context, format string, args ...any)

// SetLoggerType configures the global logger to use the specified logger type
// (ZapLogger or LogrusLogger) with default settings.
func SetLoggerType(typ LoggerType) {
	var al AbsLog
	switch typ {
	case ZapLogger:
		al = getZapLogger(defaultLogLevel, defaultEncoderType)
	case LogrusLogger:
		al = getLogrusLogger(defaultLogLevel, defaultEncoderType)
	default:
		panic(fmt.Sprintf("Logger type '%v' is not supported", typ))
	}
	SetLogger(al)
}

// SetLogger sets the provided AbsLog instance as the global logger,
// updating all global logging function variables.
func SetLogger(logger AbsLog) {
	// Debug
	Debug = logger.Debug
	Debugf = logger.Debugf
	DebugCtx = logCtx(logger.Debug)
	DebugCtxf = logCtxf(logger.Debugf)

	// Info
	Info = logger.Info
	Infof = logger.Infof
	InfoCtx = logCtx(logger.Info)
	InfoCtxf = logCtxf(logger.Infof)

	// Warn
	Warn = logger.Warn
	Warnf = logger.Warnf
	WarnCtx = logCtx(logger.Warn)
	WarnCtxf = logCtxf(logger.Warnf)

	// Error
	Error = logger.Error
	Errorf = logger.Errorf
	ErrorCtx = logCtx(logger.Error)
	ErrorCtxf = logCtxf(logger.Errorf)

	// Fatal
	Fatal = logger.Fatal
	Fatalf = logger.Fatalf
	FatalCtx = logCtx(logger.Fatal)
	FatalCtxf = logCtxf(logger.Fatalf)

	// Panic
	Panic = logger.Panic
	Panicf = logger.Panicf
	PanicCtx = logCtx(logger.Panic)
	PanicCtxf = logCtxf(logger.Panicf)
}

// getCtxValues extracts and formats context values for logging.
// It supports three formats:
// - map[string]any: formatted as "key1=value1, key2=value2"
// - []string: formatted as "item1, item2, item3"
// - string: formatted as-is
// Returns an empty string if context is nil or contains no values.
func getCtxValues(ctx context.Context) string {
	if ctx == nil || ctx.Value(contextKey) == nil {
		return ""
	}

	// Type switch to handle different context value formats
	switch ctxValues := ctx.Value(contextKey).(type) {
	case map[string]any:
		// Convert map to key=value pairs
		var pairs []string
		for k, v := range ctxValues {
			pairs = append(pairs, fmt.Sprintf("%s=%v", k, v))
		}
		return fmt.Sprintf(contextFormatTemplate, strings.Join(pairs, ", "), contextSeparator)
	case []string:
		// Join string slice with commas
		return fmt.Sprintf(contextFormatTemplate, strings.Join(ctxValues, ", "), contextSeparator)
	case string:
		// Use string value as-is
		return fmt.Sprintf(contextFormatTemplate, ctxValues, contextSeparator)
	default:
		// Unsupported type, return empty string
		return ""
	}
}

// logCtx wraps a regular log function to add context support.
// It extracts context values and prepends them to the log message.
func logCtx(log func(args ...any)) func(ctx context.Context, args ...any) {
	return func(ctx context.Context, args ...any) {
		// Extract formatted context values
		ctxValues := getCtxValues(ctx)
		if ctxValues == "" {
			// No context values, log normally
			log(args...)
		} else {
			// Prepend context values to message
			log(fmt.Sprintf("%s %s", ctxValues, fmt.Sprint(args...)))
		}
	}
}

// logCtxf wraps a formatted log function to add context support.
// It extracts context values and prepends them to the formatted log message.
func logCtxf(log func(format string, args ...any)) func(ctx context.Context, format string, args ...any) {
	return func(ctx context.Context, format string, args ...any) {
		// Extract formatted context values
		ctxValues := getCtxValues(ctx)
		if ctxValues == "" {
			// No context values, log normally
			log(format, args...)
		} else {
			// Prepend context values to formatted message
			log("%s %s", ctxValues, fmt.Sprintf(format, args...))
		}
	}
}

