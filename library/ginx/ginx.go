package ginx

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Gin struct {
	*gin.Engine
	logger gin.HandlerFunc
}

func New(bfs ...BuildFunc) *Gin {
	ret := &Gin{
		Engine: gin.New(),
	}
	for _, bf := range bfs {
		bf(ret)
	}
	return ret
}

type BuildFunc func(*Gin)

func WithLogger(logger gin.HandlerFunc) BuildFunc {
	return func(s *Gin) {
		s.logger = logger
	}
}

func (g *Gin) GetLogger() gin.HandlerFunc {
	return g.logger
}

type Logger interface {
	Info(format string, v ...interface{})
}

func DefaultLogger(logger Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		logger.Info("[FORMATTER TEST] %v | %3d | %13v | %15s | %-7s %#v\n%s",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}
}
