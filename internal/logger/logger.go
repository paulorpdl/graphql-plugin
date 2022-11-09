package logger

import "fmt"

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}

var log Logger

func SetLogger(v interface{}) {
	l, ok := v.(Logger)
	if ok {
		log = l
	}
}

func Debug(v ...interface{}) {
	log.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	log.Debug(fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	log.Info(v...)
}

func Infof(format string, v ...interface{}) {
	log.Info(fmt.Sprintf(format, v...))
}

func Warning(v ...interface{}) {
	log.Warning(v...)
}

func Error(v ...interface{}) {
	log.Error(v...)
}

func Critical(v ...interface{}) {
	log.Critical(v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}
