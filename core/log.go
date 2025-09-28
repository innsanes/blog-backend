package core

import (
	"bytes"
	"github.com/innsanes/conf"
	"github.com/innsanes/serv"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*serv.Service
	*zap.Logger
	config *LoggerConfig
}

type LoggerConfig struct {
	Json bool `conf:"json,default=false"`
}

func NewLog() *Logger {
	return &Logger{}
}

func (s *Logger) BeforeServe() (err error) {
	s.config = &LoggerConfig{}
	conf.RegisterConfWithName("log", s.config)
	return
}

func (s *Logger) Serve() (err error) {
	var encoder zapcore.Encoder
	if s.config.Json {
		encoderConfig := ZapEncoderConfigJson()
		encoderConfig.EncodeCaller = LoggerCallerEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoderConfig := ZapEncoderConfigConsole()
		encoderConfig.EncodeCaller = LoggerCallerEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	zapCore := zapcore.NewCore(encoder, os.Stdout, zapcore.DebugLevel)
	zapLogger := zap.New(zapCore, zap.AddCaller())
	s.Logger = zapLogger
	return
}

var loggerPathBufferPool *sync.Pool

func init() {
	loggerPathBufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

func ZapEncoderConfigConsole() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "message",
		StacktraceKey:    "stacktrace",
		EncodeTime:       zapcore.RFC3339TimeEncoder,
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		LineEnding:       "",
		ConsoleSeparator: " ",
	}

}

func ZapEncoderConfigJson() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "message",
		StacktraceKey:    "stacktrace",
		EncodeTime:       zapcore.RFC3339TimeEncoder,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeDuration:   zapcore.MillisDurationEncoder,
		LineEnding:       "",
		ConsoleSeparator: "",
	}
}

func LoggerCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	if !caller.Defined {
		enc.AppendString("undefined")
		return
	}
	enc.AppendString(LoggerCallerEncoderTrimmedPath(caller.File, caller.Line))
}

func ContainCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	if !caller.Defined {
		enc.AppendString("undefined")
	}
	for i := 1; i < 20; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if strings.Contains(file, "blog-backend/services") {
			enc.AppendString(LoggerCallerEncoderTrimmedPath(file, line))
			return
		}
	}
	enc.AppendString(caller.TrimmedPath())
}

func LoggerCallerEncoderTrimmedPath(file string, line int) string {
	projectFilepath := file
	index := strings.Index(file, "blog-backend")
	if index != -1 {
		projectFilepath = file[index:]
	}
	buf := loggerPathBufferPool.Get().(*bytes.Buffer)
	buf.WriteString(projectFilepath)
	buf.WriteByte(':')
	buf.WriteString(strconv.FormatInt(int64(line), 10))
	path := buf.String()
	buf.Reset()
	loggerPathBufferPool.Put(buf)
	return path
}
