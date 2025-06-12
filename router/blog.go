package router

import (
	"blog-backend/handler/blog/handler"
	"github.com/gin-gonic/gin"
)

func RegisterBlog(group *gin.RouterGroup) {
	blogGroup := group.Group("/blog")
	blogGroup.GET(":id", handler.Get)
	blogGroup.GET("/list", handler.List)
}

func RegisterBlogAuth(group *gin.RouterGroup) {
	blogGroup := group.Group("/blog")
	blogGroup.POST("/create", handler.Create)
	blogGroup.PUT("/update", handler.Update)
	blogGroup.DELETE("/delete", handler.Delete)
}
