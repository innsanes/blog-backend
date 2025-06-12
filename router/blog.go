package router

import (
	"blog-backend/global"
	"blog-backend/handler/blog/service"
	"github.com/gin-gonic/gin"
)

func init() {
	global.BlogServer.RegisterRouter(RegisterBlog)
}

func RegisterBlog(group *gin.RouterGroup) {
	blogGroup := group.Group("/blog")
	blogGroup.GET(":id", service.Get)
	blogGroup.GET("/list", service.List)
	blogGroup.POST("/create", service.Create)
	blogGroup.PUT("/update", service.Update)
	blogGroup.DELETE("/delete", service.Delete)
}
