package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type VodkaLog struct {
	name   string
	logger *zap.Logger
}

func NewVodkaLog(name string, isJsonFormat bool) *VodkaLog {
	var encoder zapcore.Encoder
	if isJsonFormat {
		encoderConfig := ZapEncoderConfigJson()
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoderConfig := ZapEncoderConfigConsole()
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	zapCore := zapcore.NewCore(encoder, os.Stdout, zapcore.DebugLevel)
	zapLogger := zap.New(zapCore)

	return &VodkaLog{
		name:   name,
		logger: zapLogger,
	}
}

func (s *VodkaLog) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestFields := buildRequestFields(c)
		c.Next()
		responseFields := buildResponseFields(c, start)

		// 合并所有字段
		allFields := append(requestFields, responseFields...)

		// 根据状态码选择日志级别
		statusCode := c.Writer.Status()
		switch {
		case statusCode >= 500:
			s.logger.Error("Vodka "+s.name+" Server Error", allFields...)
		case statusCode >= 400:
			s.logger.Warn("Vodka "+s.name+" Client Error", allFields...)
		case statusCode >= 300:
			s.logger.Info("Vodka "+s.name+" Redirect", allFields...)
		default:
			s.logger.Info("Vodka "+s.name+" Request", allFields...)
		}
	}
}

// buildRequestFields 构建请求日志字段
func buildRequestFields(c *gin.Context) []zap.Field {
	fields := []zap.Field{
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.String("client_ip", c.ClientIP()),
		zap.String("proto", c.Request.Proto),
		//zap.String("user_agent", c.Request.UserAgent()),
		//zap.String("referer", c.Request.Referer()),
		//zap.Any("headers", c.Request.Header),
		//zap.String("body", "request body logged"),
	}
	return fields
}

// buildResponseFields 构建响应日志字段
func buildResponseFields(c *gin.Context, start time.Time) []zap.Field {
	fields := []zap.Field{
		zap.Int("status", c.Writer.Status()),
		zap.Int("body_size", c.Writer.Size()),
		zap.Int64("latency", time.Since(start).Milliseconds()),
		//zap.String("response", "response body logged"),
		//zap.String("errors", c.Errors.String()),
	}
	return fields
}
