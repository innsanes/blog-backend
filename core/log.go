package core

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func NewLog() *Logger {
	zapLoggerEncoderConfig := zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "message",
		StacktraceKey:    "stacktrace",
		EncodeCaller:     LoggerCallerEncoder,
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
	zapLogger := zap.New(zapCore, zap.AddCaller(), zap.AddCallerSkip(1))

	return &Logger{
		Logger: zapLogger,
	}
}

func (s *Logger) Info(format string, v ...interface{}) {
	s.Logger.Info(fmt.Sprintf(format, v...))
}

func (s *Logger) Error(format string, v ...interface{}) {
	s.Logger.Error(fmt.Sprintf(format, v...))
}

func (s *Logger) Warn(format string, v ...interface{}) {
	s.Logger.Warn(fmt.Sprintf(format, v...))
}

func (s *Logger) Debug(format string, v ...interface{}) {
	s.Logger.Debug(fmt.Sprintf(format, v...))
}

func (s *Logger) Panic(format string, v ...interface{}) {
	s.Logger.Panic(fmt.Sprintf(format, v...))
}

var loggerPathBufferPool *sync.Pool

func init() {
	loggerPathBufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
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
	return
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
