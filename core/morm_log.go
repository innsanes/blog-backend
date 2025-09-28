package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

type MOrmLogger struct {
	*zap.Logger
	config logger.Config
}

func NewMOrmLogger(isJsonFormat bool) *MOrmLogger {
	var encoder zapcore.Encoder
	if isJsonFormat {
		encoderConfig := ZapEncoderConfigJson()
		encoderConfig.EncodeCaller = ContainCallerEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoderConfig := ZapEncoderConfigConsole()
		encoderConfig.EncodeCaller = ContainCallerEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	zapCore := zapcore.NewCore(encoder, os.Stdout, zapcore.DebugLevel)
	zapLogger := zap.New(zapCore, zap.AddCaller(), zap.AddCallerSkip(3))

	config := logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		Colorful:                  false,
		IgnoreRecordNotFoundError: false,
		ParameterizedQueries:      false,
		LogLevel:                  logger.Info,
	}

	return &MOrmLogger{
		Logger: zapLogger,
		config: config,
	}
}

func (l *MOrmLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.config.LogLevel = level
	return l
}

func (l *MOrmLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logger.Info {
		l.Logger.Info(fmt.Sprintf(msg, data...))
	}
}

func (l *MOrmLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logger.Warn {
		l.Logger.Warn(fmt.Sprintf(msg, data...))
	}
}

func (l *MOrmLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logger.Error {
		l.Logger.Error(fmt.Sprintf(msg, data...))
	}
}

func (l *MOrmLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.config.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	// 构建日志字段
	fields := []zap.Field{
		zap.String("sql", sql),
		zap.Int64("rows", rows),
		zap.Int64("elapsed", elapsed.Milliseconds()),
	}

	// 根据错误情况记录不同级别的日志
	switch {
	case err != nil && l.config.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.config.IgnoreRecordNotFoundError):
		l.Logger.Error("MOrm", append(fields, zap.Error(err))...)
	case elapsed > l.config.SlowThreshold && l.config.SlowThreshold != 0 && l.config.LogLevel >= logger.Warn:
		l.Logger.Warn("MOrm Slow SQL", fields...)
	case l.config.LogLevel >= logger.Info:
		l.Logger.Info("MOrm", fields...)
	}
}
