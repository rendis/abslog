package main

import (
	"context"
	"fmt"
	"github.com/rendis/abslog/v3"
	"github.com/rendis/abslog/v3/example/test"
)

func main() {
	// Default
	abslog.Debug("Default (Zap) Debug global")
	abslog.Info("Default (Zap) Info global")
	abslog.Warn("Default (Zap) Warn global")
	abslog.Error("Default (Zap) Error global")
	fmt.Println()

	// Default With Context
	var ctx = context.Background()
	ctxValuesMap := map[string]interface{}{
		"id":   "1234567",
		"name": "John Doe",
		"age":  30,
	}
	ctx = context.WithValue(ctx, abslog.GetCtxKey(), ctxValuesMap)

	abslog.DebugCtx(ctx, "Default (Zap) Debug with context")
	abslog.InfoCtx(ctx, "Default (Zap) Info with context")
	abslog.WarnCtx(ctx, "Default (Zap) Warn with context")
	abslog.ErrorCtx(ctx, "Default (Zap) Error with context")
	fmt.Println()

	// Print from other package
	otherpackage.PrintFromOtherPackage()

	// Set logger to Logrus by Type
	abslog.SetLoggerType(abslog.LogrusLogger)
	abslog.Debug("Set By Type Logrus Debug global")
	abslog.Info("Set By Type Logrus Info global")
	abslog.Warn("Set By Type Logrus Warn global")
	abslog.Error("Set By Type Logrus Error global")
	fmt.Println()

	// Logrus With Context
	ctx = context.Background()
	ctxValuesList := []string{
		"id: 1234567",
		"name: John Doe",
		"age: 30",
	}
	ctx = context.WithValue(ctx, abslog.GetCtxKey(), ctxValuesList)

	abslog.DebugCtx(ctx, "Logrus Debug with context")
	abslog.InfoCtx(ctx, "Logrus Info with context")
	abslog.WarnCtx(ctx, "Logrus Warn with context")
	abslog.ErrorCtx(ctx, "Logrus Error with context")
	fmt.Println()

	// Set logger to Zap by Type
	abslog.SetLoggerType(abslog.ZapLogger)
	abslog.Debug("Set By Type Zap Debug global")
	abslog.Info("Set By Type Zap Info global")
	abslog.Warn("Set By Type Zap Warn global")
	abslog.Error("Set By Type Zap Error global")
	fmt.Println()

	// Zap With Context
	ctx = context.Background()
	ctxValuesStr := "id: 1234567, name: John Doe, age: 30"
	ctx = context.WithValue(ctx, abslog.GetCtxKey(), ctxValuesStr)

	abslog.DebugCtx(ctx, "Zap Debug with context")
	abslog.InfoCtx(ctx, "Zap Info with context")
	abslog.WarnCtx(ctx, "Zap Warn with context")
	abslog.ErrorCtx(ctx, "Zap Error with context")
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
