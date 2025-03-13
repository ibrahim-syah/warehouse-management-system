package loggerutils

import "github.com/sirupsen/logrus"

type logrusEntry struct {
	entry *logrus.Entry
}

func (l *logrusEntry) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l *logrusEntry) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l *logrusEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logrusEntry) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *logrusEntry) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *logrusEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logrusEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *logrusEntry) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *logrusEntry) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logrusEntry) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l *logrusEntry) WithField(key string, value interface{}) (entry Logger) {
	entry = &logrusEntry{l.entry.WithField(key, value)}
	return
}

func (l *logrusEntry) WithFields(args map[string]interface{}) (entry Logger) {
	entry = &logrusEntry{l.entry.WithFields(args)}
	return
}
