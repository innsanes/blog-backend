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
	blogGroup.GET("/draft/:id", service.GetBlogDraft)

	draftGroup := group.Group("/blog-draft")
	draftGroup.GET(":id", service.GetDraft)
	draftGroup.GET("/list", service.ListDraft)
	draftGroup.POST("/create", service.CreateDraft)
	draftGroup.POST("/create-blog", service.CreateBlogFromDraft)
	draftGroup.PUT("/update", service.UpdateDraft)
}
