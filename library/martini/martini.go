package martini

import (
	"github.com/gin-gonic/gin"
	"time"
)

type M struct {
	*gin.Engine
	logger gin.HandlerFunc
}

func New(bfs ...BuildFunc) *M {
	ret := &M{
		Engine: gin.New(),
		logger: gin.Logger(),
	}
	for _, bf := range bfs {
		bf(ret)
	}
	return ret
}

type BuildFunc func(*M)

func WithLogHandler(logger LogHandler) BuildFunc {
	return func(s *M) {
		s.logger = defaultLogger(logger)
	}
}

type LogHandler func(param gin.LogFormatterParams)

func (g *M) Logger() gin.HandlerFunc {
	return g.logger
}

func defaultLogger(logger LogHandler) gin.HandlerFunc {
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
		logger(param)
	}
}
