package transport

type Logger interface {
	Debug(message string)
	Debugf(format string, v ...any)
	Error(message string)
	Errorf(format string, v ...any)
	Info(message string)
	Infof(format string, v ...any)
	Warn(message string)
	Warnf(format string, v ...any)
}
