package main

import (
	"github.com/Mr-Punder/gophermart/internal/config"
	"github.com/Mr-Punder/gophermart/internal/logger"
)

func main() {
	conf := config.New()
	log, err := logger.NewZapLogger(conf.LogLevel, conf.LogOutputPath)
	if err != nil {
		panic(err)
	}
	log.Info("Initializwd logger")

}
