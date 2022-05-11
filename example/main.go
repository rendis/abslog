package main

import (
	"github.com/rendis/abslog"
)

var log = GetLogger()

func GetLogger() *abslog.AbsLog {
	return abslog.GetAbsLogBuilder().LogLevel(abslog.InfoLevel).Build()
}

func main() {
	log.Debug("Debug not logged")
	log.Info("Info logged")
	log.Warn("Warn logged")
	log.Error("Error logged with stacktrace")
	doProcess()
	log.Fatal("Fatal logged with stacktrace and exit")
	log.Info("Info not logged")
}
