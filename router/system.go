package router

import (
	"blog-backend/services/handler"
	"net/http/pprof"

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

func RegisterPProf(group *gin.RouterGroup) {
	pprofGroup := group.Group("/pprof")
	{
		pprofGroup.GET("/index", gin.WrapF(pprof.Index))
		pprofGroup.GET("/profile", gin.WrapF(pprof.Profile))
		pprofGroup.GET("/heap", gin.WrapF(pprof.Handler("heap").ServeHTTP))
		pprofGroup.GET("/goroutine", gin.WrapF(pprof.Handler("goroutine").ServeHTTP))
		pprofGroup.GET("/block", gin.WrapF(pprof.Handler("block").ServeHTTP))
		pprofGroup.GET("/mutex", gin.WrapF(pprof.Handler("mutex").ServeHTTP))
		pprofGroup.GET("/allocs", gin.WrapF(pprof.Handler("allocs").ServeHTTP))
		pprofGroup.GET("/threadcreate", gin.WrapF(pprof.Handler("threadcreate").ServeHTTP))
		pprofGroup.GET("/symbol", gin.WrapF(pprof.Symbol))
		pprofGroup.GET("/trace", gin.WrapF(pprof.Trace))
		pprofGroup.GET("/cmdline", gin.WrapF(pprof.Cmdline))
	}
}
