package log

import (
	"log"
	"runtime"
)

func log_with_level(level string, message string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	}

	log.Printf("[%s] %s:%d - %s\n", level, file, line, message)
}

func Debug[T any](message string, err T) {
	log_with_level("DEBUG", message)
}

func Info[T any](message string, err T) {
	log_with_level("INFO", message)
}

func Warn[T any](message string, err T) {
	log_with_level("WARN", message)
}

func Error[T any](message string, err T) {
	log_with_level("ERROR", message)
}

func Fatal[T any](message string, err T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	log.Fatalf("[FATAL] %s:%d - %v\n", file, line, message)
}

