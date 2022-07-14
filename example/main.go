package main

import (
	"fmt"
	"github.com/rendis/abslog"
)

func main() {
	// Default
	abslog.Debug("Default (Zap) Debug global")
	abslog.Info("Default (Zap) Info global")
	abslog.Warn("Default (Zap) Warn global")
	abslog.Error("Default (Zap) Error global")
	fmt.Println()

	// Set logger to Logrus by Type
	abslog.SetLoggerType(abslog.LogrusLogger)
	abslog.Debug("Set By Type Logrus Debug global")
	abslog.Info("Set By Type Logrus Info global")
	abslog.Warn("Set By Type Logrus Warn global")
	abslog.Error("Set By Type Logrus Error global")
	fmt.Println()

	// Set logger to Zap by Type
	abslog.SetLoggerType(abslog.ZapLogger)
	abslog.Debug("Set By Type Zap Debug global")
	abslog.Info("Set By Type Zap Info global")
	abslog.Warn("Set By Type Zap Warn global")
	abslog.Error("Set By Type Zap Error global")
	fmt.Println()

	// Set logger to Logrus by Builder
	abslog.GetAbsLogBuilder().
		LoggerType(abslog.LogrusLogger).
		LogLevel(abslog.InfoLevel).
		BuildAndSetAsGlobal()
	abslog.Debug("Set By Builder Logrus Debug global")
	abslog.Info("Set By Builder Logrus Info global")
	abslog.Warn("Set By Builder Logrus Warn global")
	abslog.Error("Set By Builder Logrus Error global")
	fmt.Println()

	// Set logger to Zap by Builder
	abslog.GetAbsLogBuilder().
		LoggerType(abslog.ZapLogger).
		LogLevel(abslog.InfoLevel).
		BuildAndSetAsGlobal()
	abslog.Debug("Set By Builder Zap Debug global")
	abslog.Info("Set By Builder Zap Info global")
	abslog.Warn("Set By Builder Zap Warn global")
	abslog.Error("Set By Builder Zap Error global")
	fmt.Println()
}
