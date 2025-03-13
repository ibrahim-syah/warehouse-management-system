package loggerutils

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type logrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger() *logrusLogger {
	log := logrus.New()

	env := viper.GetString("ENV")
	if env != "PROD" {
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
		})
		log.SetLevel(logrus.TraceLevel)
		log.Out = os.Stdout
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
		})
		log.SetLevel(logrus.InfoLevel)

		today := time.Now().Format("2006-01-02")
		filename := fmt.Sprintf("%s.log", today)
		filepath := fmt.Sprintf("logs/%s", filename)
		file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	}

	return &logrusLogger{
		log: log,
	}
}

func (l *logrusLogger) Debug(args ...any) {
	l.log.Debug(args...)
}

func (l *logrusLogger) Debugf(format string, args ...any) {
	l.log.Debugf(format, args...)
}

func (l *logrusLogger) Info(args ...any) {
	l.log.Info(args...)
}

func (l *logrusLogger) Infof(format string, args ...any) {
	l.log.Infof(format, args...)
}

func (l *logrusLogger) Warn(args ...any) {
	l.log.Warn(args...)
}

func (l *logrusLogger) Warnf(format string, args ...any) {
	l.log.Warnf(format, args...)
}

func (l *logrusLogger) Error(args ...any) {
	l.log.Error(args...)
}

func (l *logrusLogger) Errorf(format string, args ...any) {
	l.log.Errorf(format, args...)
}

func (l *logrusLogger) Fatal(args ...any) {
	l.log.Fatal(args...)
}

func (l *logrusLogger) Fatalf(format string, args ...any) {
	l.log.Fatalf(format, args...)
}

func (l *logrusLogger) WithField(key string, value any) Logger {
	return &logrusEntry{
		entry: l.log.WithField(key, value),
	}
}

func (l *logrusLogger) WithFields(fields map[string]any) Logger {
	return &logrusEntry{
		entry: l.log.WithFields(fields),
	}
}
