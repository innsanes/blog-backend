package router

import (
	"blog-backend/services/handler"
	"github.com/gin-gonic/gin"
)

func RegisterEcho(group *gin.RouterGroup) {
	echoGroup := group.Group("/echo")
	echoGroup.GET("", handler.Echo)
}
