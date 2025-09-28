# abslog

[![Go Reference](https://pkg.go.dev/badge/github.com/rendis/abslog/v3.svg)](https://pkg.go.dev/github.com/rendis/abslog/v3)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.25.1-blue.svg)](https://golang.org/)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![CodeQL](https://github.com/rendis/abslog/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/rendis/abslog/actions/workflows/github-code-scanning/codeql)
[![Dependabot](https://img.shields.io/badge/Dependabot-enabled-brightgreen)](https://github.com/rendis/abslog/security/dependabot)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/rendis/abslog)
[![Contributing](https://img.shields.io/badge/contributing-guide-blue.svg)](CONTRIBUTING.md)

A modern Go logging abstraction library that provides a unified API over multiple logging backends. It ships with built-in support for popular libraries like Zap and Logrus, enabling seamless switching between them without code changes. Additionally, it allows integration of any custom logging library through its adapter pattern. Includes powerful context-aware logging for enhanced traceability in distributed systems.

## Features

- **Built-in Backends**: Ready-to-use integration with Zap and Logrus
- **Extensible Design**: Add any logging library via the LoggerAdapter interface
- **Unified API**: Consistent logging interface across all backends
- **Backend Flexibility**: Switch between supported loggers without code modifications
- **Context Logging**: Embed contextual information (e.g., transaction IDs, user data) in logs for traceability
- **Builder Pattern**: Fluent configuration API for logger setup
- **Multiple Output Formats**: Support for console and JSON encoding
- **Global Functions**: Ready-to-use global logging functions with context support

## Installation

```bash
go get github.com/rendis/abslog/v3
```

## Quick Start

```go
package main

import (
    "github.com/rendis/abslog/v3"
)

func main() {
    abslog.Info("Application started")
    abslog.Error("An error occurred")
}
```

## Usage

### Basic Logging

abslog provides global logging functions at all standard levels:

```go
abslog.Debug("Debug message")
abslog.Info("Info message")
abslog.Warn("Warning message")
abslog.Error("Error message")
abslog.Fatal("Fatal message") // Exits the program
abslog.Panic("Panic message") // Panics the program
```

Formatted logging is also supported:

```go
abslog.Infof("User %s logged in at %s", username, time.Now())
```

### Switching Backends

Change the underlying logging library without modifying your logging code:

```go
// Switch to Logrus
abslog.SetLoggerType(abslog.LogrusLogger)

// Switch back to Zap (default)
abslog.SetLoggerType(abslog.ZapLogger)
```

### Context Logging

abslog's context logging enables powerful traceability features, particularly useful in microservices and distributed systems. By embedding contextual information in the `context.Context`, you can correlate logs across request lifecycles.

#### Setting Context Values

At the start of a request or operation, set contextual data:

```go
ctx := context.Background()

// Using a map for structured context
ctxValues := map[string]any{
    "transaction_id": "txn-12345",
    "user_id":        "user-67890",
    "service":        "auth-service",
}
ctx = context.WithValue(ctx, abslog.GetCtxKey(), ctxValues)

// Or using a slice of strings
ctx = context.WithValue(ctx, abslog.GetCtxKey(), []string{"txn-12345", "user-67890"})

// Or a simple string
ctx = context.WithValue(ctx, abslog.GetCtxKey(), "txn-12345")
```

#### Logging with Context

Use context-aware logging functions to include the embedded data in your logs:

```go
abslog.InfoCtx(ctx, "Processing user authentication")
abslog.WarnCtx(ctx, "Invalid credentials provided")
abslog.ErrorCtx(ctx, "Authentication failed")
```

**Output Example:**

```textplain
[transaction_id=txn-12345, user_id=user-67890, service=auth-service] -> Processing user authentication
[transaction_id=txn-12345, user_id=user-67890, service=auth-service] -> Invalid credentials provided
```

This allows you to trace all logs related to a specific transaction or user across your entire application, making debugging and monitoring significantly easier.

### Advanced Configuration

Use the builder pattern for detailed logger configuration:

```go
// BuildAndSetAsGlobal: Creates the logger and sets it as the global logger
logger := abslog.GetAbsLogBuilder().
    LoggerType(abslog.LogrusLogger).
    LogLevel(abslog.DebugLevel).
    EncoderType(abslog.JSONEncoder).
    ContextKey("custom-key").
    BuildAndSetAsGlobal()

// Build: Creates the logger instance without setting it as global
customLogger := abslog.GetAbsLogBuilder().
    LoggerType(abslog.ZapLogger).
    LogLevel(abslog.InfoLevel).
    Build()

// Use the custom logger directly (not affecting global functions)
customLogger.Info("This uses the custom logger instance")
```

**Difference between Build and BuildAndSetAsGlobal:**

- `Build()`: Returns a configured `AbsLog` instance that you can use directly, but doesn't affect the global logging functions
- `BuildAndSetAsGlobal()`: Configures the logger and sets it as the global logger, updating all global `abslog.Info()`, `abslog.Debug()`, etc. functions to use this configuration

#### Custom Context Key

Customize the context key used for storing values:

```go
abslog.SetCtxKey("my-custom-key")
```

### Adding Custom Logging Libraries

abslog is designed to be extensible. You can integrate any logging library that provides the standard logging methods. The process involves creating a generator function and using the LoggerAdapter.

#### Implementation Steps

1. **Create a Generator Function**: Implement a function that takes `LogLevel` and `EncoderType` and returns an `AbsLog`:

    ```go
    func getCustomLogger(logLevel LogLevel, encoder EncoderType) AbsLog {
        // Create your custom logger instance
        customLogger := // ... initialize your logger

        // Configure log level
        customLogger.SetLevel(convertToCustomLevel(logLevel))

        // Configure encoding if supported
        switch encoder {
        case JSONEncoder:
            // Set JSON formatter
        case ConsoleEncoder:
            // Set console formatter
        }

        // Wrap in LoggerAdapter
        return NewLoggerAdapter(customLogger)
    }
    ```

2. **Level Conversion**: Create a helper function to convert abslog levels to your library's levels:

    ```go
    func convertToCustomLevel(logLevel LogLevel) CustomLevel {
        switch logLevel {
        case DebugLevel:
            return CustomDebug
        case InfoLevel:
            return CustomInfo
        // ... other levels
        default:
            return CustomInfo
        }
    }
    ```

3. **Use with Builder**: Set your custom generator and build the logger:

```go
logger := abslog.GetAbsLogBuilder().
    LoggerGen(getCustomLogger).
    LogLevel(abslog.DebugLevel).
    BuildAndSetAsGlobal()
```

#### Examples

See [`logrus.go`](logrus.go) and [`zap.go`](zap.go) for complete implementations of Logrus and Zap integrations. These files demonstrate:

- Logger initialization and configuration
- Level conversion functions
- Encoder setup for console and JSON output
- Proper wrapping with `NewLoggerAdapter`

The LoggerAdapter requires your logger to implement methods: `Debug/Info/Warn/Error/Fatal/Panic` and their formatted variants (`Debugf/Infof/etc.`).

## API Overview

### Global Functions

- `Debug/Info/Warn/Error/Fatal/Panic(args ...any)`
- `Debugf/Infof/Warnf/Errorf/Fatalf/Panicf(format string, args ...any)`
- `DebugCtx/InfoCtx/WarnCtx/ErrorCtx/FatalCtx/PanicCtx(ctx context.Context, args ...any)`
- `DebugCtxf/InfoCtxf/WarnCtxf/ErrorCtxf/FatalCtxf/PanicCtxf(ctx context.Context, format string, args ...any)`

### Configuration

- `SetLoggerType(LoggerType)`
- `SetLogger(AbsLog)`
- `GetAbsLogBuilder() AbsLogBuilder`

### Context Management

- `SetCtxKey(key string)`
- `GetCtxKey() string`
- `SetCtxSeparator(separator string)`

### Types

- `LoggerType`: `ZapLogger`, `LogrusLogger`
- `LogLevel`: `DebugLevel`, `InfoLevel`, `WarnLevel`, `ErrorLevel`, `FatalLevel`, `PanicLevel`
- `EncoderType`: `ConsoleEncoder`, `JSONEncoder`

## Contributing

- 🤝 [**Contributing Guide**](CONTRIBUTING.md) - How to contribute code, report issues, and help improve abslog

## License

GPL v3 - see [LICENSE](LICENSE) for details.
