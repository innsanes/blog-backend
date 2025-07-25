package router

import (
	"blog-backend/services/handler"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterEcho(group *gin.RouterGroup) {
	echoGroup := group.Group("/echo")
	echoGroup.GET("", handler.Echo)
}

func RegisterPrometheus(group *gin.RouterGroup) {
	prometheusGroup := group.Group("/metrics")
	prometheusGroup.GET("", gin.WrapH(promhttp.Handler()))
}
