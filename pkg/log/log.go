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
	if (level == "Debug") {
		log.Printf("[%s] %s:%d - %s\n", level, file, line, message)
	} else {
		log.Printf("[%s] - %s\n", level, message)
	}
}

func Debug(message string) {
	log_with_level("DEBUG", message)
}

func Info(message string) {
	log_with_level("INFO", message)
}

func Warn(message string) {
	log_with_level("WARN", message)
}

func Error(message string) {
	log_with_level("ERROR", message)
}

func Fatal(message string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	log.Fatalf("[FATAL] %s:%d - %v\n", file, line, message)
}

