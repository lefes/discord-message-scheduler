package logger

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

const (
	// DefaultLogLevel is the default log level
	DefaultLogLevel = "info"
	// DefaultLogFormat is the default log format
	DefaultLogFormat = "json"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	Log = logrus.New()

	logLevel, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevel = DefaultLogLevel
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		Log.Errorf("Invalid log level: %v", err)
		level = logrus.InfoLevel
	}

	Log.SetLevel(level)

	logFormat, ok := os.LookupEnv("LOG_FORMAT")
	if !ok {
		logFormat = DefaultLogFormat
	}

	switch logFormat {
	case "json":
		Log.SetFormatter(&logrus.JSONFormatter{})
	default:
		Log.SetFormatter(&logrus.TextFormatter{})
	}
}
