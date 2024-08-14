package core

import (
	"fmt"
	"github.com/innsanes/serv"
	"go.uber.org/zap"
)

type Log struct {
	config *zap.Config
	*serv.Service
	*zap.Logger
}

func NewLog() *Log {
	config := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	return &Log{
		config: config,
	}
}

func (s *Log) BeforeServe() (err error) {
	s.Logger, err = s.config.Build()
	return
}

func (s *Log) Info(format string, v ...interface{}) {
	s.Logger.Info(fmt.Sprintf(format, v...))
	return
}
