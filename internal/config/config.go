package config

import (
	"flag"
	"os"
)

type Config struct {
	RunAddress    string
	LogLevel      string
	LogOutputPath string
	LogErrorPath  string
	AccrualPath   string
	DBstring      string
}

// New from environment and consol parameters
func New() *Config {
	var (
		flagRunAddr, logLevel, logOutputPath, accuralPath, logErrortPath, dbString string
	)

	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "addres and port to run server")
	flag.StringVar(&logLevel, "l", "info", "level of logging")
	flag.StringVar(&logOutputPath, "lp", "stdout", "log output path")
	flag.StringVar(&logErrortPath, "le", "stderr", "log error output path")
	flag.StringVar(&accuralPath, "r", "/tmp/metrics-db.json", "storage filename")
	flag.StringVar(&dbString, "d", "", "databese opening string")
	// host=localhost user=metrics password=metrics_password dbname=metrics

	flag.Parse()

	if envAddrs, ok := os.LookupEnv("RUN_ADDRESS"); ok {

		flagRunAddr = envAddrs
	}
	if envLogLevel, ok := os.LookupEnv("LOGLEVEL"); ok {

		logLevel = envLogLevel
	}
	if envLogOutputPath, ok := os.LookupEnv("LOGPATH"); ok {

		logOutputPath = envLogOutputPath
	}
	if envLogErrorPath, ok := os.LookupEnv("LOGERROR"); ok {

		logErrortPath = envLogErrorPath
	}

	if envAccuralPath, ok := os.LookupEnv("ACCRUAL_SYSTEM_ADDRESS"); ok {

		accuralPath = envAccuralPath
	}
	if envDBstring, ok := os.LookupEnv("DATABASE_URI"); ok {
		dbString = envDBstring
	}

	return &Config{
		RunAddress:    flagRunAddr,
		LogLevel:      logLevel,
		LogOutputPath: logOutputPath,
		LogErrorPath:  logErrortPath,
		AccrualPath:   accuralPath,
		DBstring:      dbString,
	}
}
