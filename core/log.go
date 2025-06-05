package core

import (
	"fmt"
	"github.com/innsanes/serv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	*serv.Service
	*zap.Logger
	conf   Confer
	config *LogConfig
}

type LogConfig struct {
	Level string `conf:"level,default=debug"`
}

func NewLog() *Logger {
	zapLoggerEncoderConfig := zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "message",
		StacktraceKey:    "stacktrace",
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeTime:       zapcore.RFC3339TimeEncoder,
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		LineEnding:       "",
		ConsoleSeparator: " ",
	}

	zapCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapLoggerEncoderConfig),
		os.Stdout,
		zapcore.DebugLevel,
	)
	zapLogger := zap.New(zapCore, zap.AddCaller())

	return &Logger{
		Logger: zapLogger,
		//conf:   conf,
		//config: &LogConfig{},
	}
}

func (s *Logger) BeforeServe() (err error) {
	//s.conf.RegisterConfWithName("log", s.config)
	return
}

func (s *Logger) Serve() (err error) {
	return
}

func (s *Logger) Info(format string, v ...interface{}) {
	//s.Logger.Info(fmt.Sprintf(format, v...))
	s.Logger.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(format, v...))
	return
}

func (s *Logger) Error(format string, v ...interface{}) {
	s.Logger.Error(fmt.Sprintf(format, v...))
	return
}

func (s *Logger) Warn(format string, v ...interface{}) {
	s.Logger.Warn(fmt.Sprintf(format, v...))
	return
}

func (s *Logger) Debug(format string, v ...interface{}) {
	s.Logger.Debug(fmt.Sprintf(format, v...))
	return
}
