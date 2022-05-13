package main

import (
	"fmt"
	"github.com/rendis/abslog"
)

func GetZapLogger() *abslog.AbsLog {
	return abslog.GetAbsLogBuilder().
		LoggerType(abslog.ZapLogger).
		LogLevel(abslog.DebugLevel).
		Build()
}

func GetLogrusLogger() *abslog.AbsLog {
	return abslog.GetAbsLogBuilder().
		LoggerType(abslog.LogrusLogger).
		LogLevel(abslog.InfoLevel).
		Build()
}

func main() {
	var log = GetZapLogger()
	useLog(log, "Zap logger")

	fmt.Println()

	log = GetLogrusLogger()
	useLog(log, "Logrus logger")
}

func useLog(log *abslog.AbsLog, logType string) {
	log.Debug("Debug logged - ", logType)
	log.Info("Info logged - ", logType)
	log.Warn("Warn logged - ", logType)
	log.Error("Error logged with stacktrace - ", logType)
}
