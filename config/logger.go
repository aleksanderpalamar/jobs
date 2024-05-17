package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG:", logger.Flags()),
		info:    log.New(writer, "INFO:", logger.Flags()),
		warning: log.New(writer, "WARNING:", logger.Flags()),
		err:     log.New(writer, "ERROR:", logger.Flags()),
		writer:  writer,
	}
}

// Create Non-Formatted Logs
func (l *Logger) Debug(args ...interface{}) {
	l.debug.Println(args...)
}
func (l *Logger) Info(args ...interface{}) {
	l.info.Println(args...)
}
func (l *Logger) Warning(args ...interface{}) {
	l.warning.Println(args...)
}
func (l *Logger) Error(args ...interface{}) {
	l.err.Println(args...)
}

// Create Format Enabled Logs
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.debug.Printf(format, args...)
}
func (l *Logger) Infof(format string, args ...interface{}) {
	l.info.Printf(format, args...)
}
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.warning.Printf(format, args...)
}
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.err.Printf(format, args...)
}
