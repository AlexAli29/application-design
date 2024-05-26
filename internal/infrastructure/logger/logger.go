package logger

import (
	"fmt"
	"log"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type Logger struct {
	logger   *log.Logger
	minLevel LogLevel
}

func NewLogger(minLevel LogLevel) *Logger {
	return &Logger{logger: log.Default(), minLevel: minLevel}
}

func (l *Logger) logf(level LogLevel, prefix, format string, v ...any) {
	if level >= l.minLevel {
		msg := fmt.Sprintf(format, v...)
		l.logger.Printf("[%s]: %s\n", prefix, msg)
	}
}

func (l *Logger) log(level LogLevel, prefix, msg string) {
	if level >= l.minLevel {

		l.logger.Printf("[%s]: %s\n", prefix, msg)
	}
}

func (l *Logger) Debugf(format string, v ...any) {
	l.logf(DEBUG, "DEBUG", format, v...)
}

func (l *Logger) Infof(format string, v ...any) {
	l.logf(INFO, "INFO", format, v...)
}

func (l *Logger) Warnf(format string, v ...any) {
	l.logf(WARN, "WARN", format, v...)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.logf(ERROR, "ERROR", format, v...)
}

func (l *Logger) Debug(message string) {
	l.log(DEBUG, "DEBUG", message)
}

func (l *Logger) Info(message string) {
	l.log(INFO, "INFO", message)
}

func (l *Logger) Warn(message string) {
	l.log(WARN, "WARN", message)
}

func (l *Logger) Error(message string) {
	l.log(ERROR, "ERROR", message)
}
