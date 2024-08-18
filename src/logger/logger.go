package logger

import (
	"github.com/natefinch/lumberjack"
	"log"
	"os"
)

// NewLoggers to log and manage log files
func NewLoggers() (*log.Logger, *log.Logger) {
	// Check if log directories exist, create them if not
	logDirs := []string{"log/application", "log/error"}
	for _, dir := range logDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
	}

	// Set up logger for application logs
	appLog := &lumberjack.Logger{
		Filename:   "log/application/app.log",
		MaxSize:    500, // mB
		MaxBackups: 3,
		MaxAge:     7, // days
		Compress:   true,
	}
	appLogger := log.New(appLog, "", log.LstdFlags)

	// Set up logger for error logs
	errorLog := &lumberjack.Logger{
		Filename:   "log/error/error.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
	}
	errorLogger := log.New(errorLog, "", log.LstdFlags)

	return appLogger, errorLogger
}
