package middleware

import (
	"blog-backend/global/prom"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		// 记录请求持续时间
		duration := time.Since(start).Seconds()
		prom.HttpRequestDuration.WithLabelValues(c.Request.Method, c.FullPath()).Observe(duration)

		// 记录请求计数
		status := strconv.Itoa(c.Writer.Status())
		prom.HttpRequestCount.WithLabelValues(c.Request.Method, c.FullPath(), status).Inc()
	}
}
