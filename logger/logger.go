package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel = "LOG_LEVEL"
)

var (
	log logger
)

type BookStoreLogger interface {
	Printf(format string, v ...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "json",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func GetLogger() BookStoreLogger {
	return log
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	_ = log.log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}
	log.log.Error(msg, tags...)
	_ = log.log.Sync()
}
