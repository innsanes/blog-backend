package router

import (
	"blog-backend/services/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBlog(group *gin.RouterGroup) {
	blogGroup := group.Group("/blog")
	blogGroup.GET(":id", handler.BlogGet)
	blogGroup.GET("", handler.BlogList)
}

func RegisterBlogAuth(group *gin.RouterGroup) {
	blogGroup := group.Group("/blog")
	blogGroup.POST("", handler.BlogCreate)
	blogGroup.PUT(":id", handler.BlogUpdate)
	blogGroup.DELETE(":id", handler.BlogDelete)
}
