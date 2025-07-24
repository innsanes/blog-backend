package router

import (
	"blog-backend/services/blog/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBlog(group *gin.RouterGroup) {
	blogGroup := group.Group("/blog")
	blogGroup.GET(":id", handler.Get)
	blogGroup.GET("", handler.List)
}

func RegisterBlogAuth(group *gin.RouterGroup) {
	blogGroup := group.Group("/blog")
	blogGroup.POST("", handler.Create)
	blogGroup.PUT(":id", handler.Update)
	blogGroup.DELETE(":id", handler.Delete)
}
