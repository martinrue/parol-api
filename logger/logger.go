package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// Logger handles system and request logging.
type Logger struct {
	logger *log.Logger
}

// Print writes a message and any args to the logger.
func (l *Logger) Print(rid string, message string, args ...interface{}) {
	l.logger.Printf(l.withContext(rid, l.callerName(), message), args...)
}

// System writes a message without rid and caller to the logger.
func (l *Logger) System(message string, args ...interface{}) {
	l.logger.Printf(message, args...)
}

func (l *Logger) withContext(rid string, caller string, message string) string {
	return fmt.Sprintf("<%s> %s: %s\n", rid, caller, message)
}

func (l *Logger) callerName() string {
	pc, _, _, _ := runtime.Caller(2)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	return parts[len(parts)-1]
}

// New creates a new logger.
func New(prefix string) *Logger {
	return &Logger{
		logger: log.New(os.Stderr, prefix+" ", log.LstdFlags),
	}
}
