package abslog

// LoggerAdapter adapts any logger that implements the basic logging methods
// to the AbsLog interface. This provides a consistent abstraction layer
// while handling type conversions.
type LoggerAdapter struct {
	logger interface {
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
}

// NewLoggerAdapter creates a new LoggerAdapter wrapping the provided logger.
func NewLoggerAdapter(logger interface {
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
}) AbsLog {
	return &LoggerAdapter{logger: logger}
}

// Debug logs a message at debug level.
func (a *LoggerAdapter) Debug(args ...any) {
	a.logger.Debug(args...)
}

// Debugf logs a formatted message at debug level.
func (a *LoggerAdapter) Debugf(format string, args ...any) {
	a.logger.Debugf(format, args...)
}

// Info logs a message at info level.
func (a *LoggerAdapter) Info(args ...any) {
	a.logger.Info(args...)
}

// Infof logs a formatted message at info level.
func (a *LoggerAdapter) Infof(format string, args ...any) {
	a.logger.Infof(format, args...)
}

// Warn logs a message at warn level.
func (a *LoggerAdapter) Warn(args ...any) {
	a.logger.Warn(args...)
}

// Warnf logs a formatted message at warn level.
func (a *LoggerAdapter) Warnf(format string, args ...any) {
	a.logger.Warnf(format, args...)
}

// Error logs a message at error level.
func (a *LoggerAdapter) Error(args ...any) {
	a.logger.Error(args...)
}

// Errorf logs a formatted message at error level.
func (a *LoggerAdapter) Errorf(format string, args ...any) {
	a.logger.Errorf(format, args...)
}

// Fatal logs a message at fatal level and exits the program.
func (a *LoggerAdapter) Fatal(args ...any) {
	a.logger.Fatal(args...)
}

// Fatalf logs a formatted message at fatal level and exits the program.
func (a *LoggerAdapter) Fatalf(format string, args ...any) {
	a.logger.Fatalf(format, args...)
}

// Panic logs a message at panic level and panics.
func (a *LoggerAdapter) Panic(args ...any) {
	a.logger.Panic(args...)
}

// Panicf logs a formatted message at panic level and panics.
func (a *LoggerAdapter) Panicf(format string, args ...any) {
	a.logger.Panicf(format, args...)
}
