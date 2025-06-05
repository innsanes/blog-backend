package router

import (
	"blog-backend/global"
	"blog-backend/handler/image/service"
	"github.com/gin-gonic/gin"
)

func init() {
	global.BlogServer.RegisterRouter(RegisterImage)
}

func RegisterImage(group *gin.RouterGroup) {
	imageGroup := group.Group("/image")

	imageGroup.GET(":name", service.Get)
	imageGroup.POST("upload", service.Upload)
}
