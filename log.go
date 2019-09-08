package main

import (
	"github.com/kasraf/golog"
	"os"
)

var logFile *os.File
var logger *GoLog.Logger

func initLogger() {
	var err error

	logFile, err = os.Create("dev.log")

	if err != nil {
		panic(err)
	}

	GoLog.Init(logFile)
	logger = GoLog.GetLogger()
}

func closeLogger() {
	handle(logFile.Close())
}

func handle(err error) {
	if err != nil {
		logger.Error("Tried to unwrap non-nil error: ", err)
		panic(err)
	}
}
