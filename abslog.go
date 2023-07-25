package abslog

import (
	"context"
	"fmt"
	"runtime"
	"strings"
)

var contextKey = "abslog"
var contextSeparator = " -> "

type AbsLog interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
}

func init() {
	fmt.Println("init abslog with default logger type (zap)")
	SetLoggerType(defaultLoggerType)
}

func SetCtxKey(key string) {
	contextKey = key
}

func GetCtxKey() string {
	return contextKey
}

func SetCtxSeparator(separator string) {
	contextSeparator = separator
}

// Debug logs a message at level Debug on the standard logger.
var Debug func(args ...interface{})
var DebugCtx func(ctx context.Context, args ...interface{})
var Debugf func(format string, args ...interface{})
var DebugCtxf func(ctx context.Context, format string, args ...interface{})

// Info logs a message at level Info on the standard logger.
var Info func(args ...interface{})
var InfoCtx func(ctx context.Context, args ...interface{})
var Infof func(format string, args ...interface{})
var InfoCtxf func(ctx context.Context, format string, args ...interface{})

// Warn logs a message at level Warn on the standard logger.
var Warn func(args ...interface{})
var WarnCtx func(ctx context.Context, args ...interface{})
var Warnf func(format string, args ...interface{})
var WarnCtxf func(ctx context.Context, format string, args ...interface{})

// Error logs a message at level Error on the standard logger.
var Error func(args ...interface{})
var ErrorCtx func(ctx context.Context, args ...interface{})
var Errorf func(format string, args ...interface{})
var ErrorCtxf func(ctx context.Context, format string, args ...interface{})

// Fatal logs a message at level Fatal on the standard logger.
var Fatal func(args ...interface{})
var FatalCtx func(ctx context.Context, args ...interface{})
var Fatalf func(format string, args ...interface{})
var FatalCtxf func(ctx context.Context, format string, args ...interface{})

// Panic logs a message at level Panic on the standard logger.
var Panic func(args ...interface{})
var PanicCtx func(ctx context.Context, args ...interface{})
var Panicf func(format string, args ...interface{})
var PanicCtxf func(ctx context.Context, format string, args ...interface{})

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

func getCtxValues(ctx context.Context) string {
	if ctx == nil || ctx.Value(contextKey) == nil {
		return ""
	}

	switch ctxValues := ctx.Value(contextKey).(type) {
	case map[string]interface{}:
		var pairs []string
		for k, v := range ctxValues {
			pairs = append(pairs, fmt.Sprintf("%s=%v", k, v))
		}
		return fmt.Sprintf("[%s]%s", strings.Join(pairs, ", "), contextSeparator)
	case []string:
		return fmt.Sprintf("[%s]%s", strings.Join(ctxValues, ", "), contextSeparator)
	case string:
		return fmt.Sprintf("[%s]%s", ctxValues, contextSeparator)
	default:
		return ""
	}
}

func logCtx(log func(args ...interface{})) func(ctx context.Context, args ...interface{}) {
	return func(ctx context.Context, args ...interface{}) {
		callerInfo := getCallerInfo()
		ctxValues := getCtxValues(ctx)
		log(fmt.Sprintf("%s %s %s", callerInfo, ctxValues, fmt.Sprint(args...)))
	}
}

func logCtxf(log func(format string, args ...interface{})) func(ctx context.Context, format string, args ...interface{}) {
	return func(ctx context.Context, format string, args ...interface{}) {
		callerInfo := getCallerInfo()
		ctxValues := getCtxValues(ctx)
		log(fmt.Sprintf("%s %s %s", callerInfo, ctxValues, fmt.Sprintf(format, args...)))
	}
}

func getCallerInfo() string {
	pc, fileName, lineNumber, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	funcName := runtime.FuncForPC(pc).Name()

	parts := strings.Split(fileName, "/")
	fileName = strings.Join(parts[:len(parts)-1], "/")

	return fmt.Sprintf("[%s/%s:%d]", fileName, funcName, lineNumber)
}
