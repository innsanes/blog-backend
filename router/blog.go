package router

import (
	"blog-backend/global"
	"blog-backend/handler/blog/service"
	"blog-backend/handler/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	global.Gin.RegisterRouter(RegisterBlog)
}

func RegisterBlog(group *gin.RouterGroup) {
	blogGroup := group.Group("/blog")
	blogGroup.GET(":id", service.Get)
	blogGroup.GET("/list", service.List)
	blogGroup.Use(middleware.TokenCheck()).POST("/create", service.Create)
	blogGroup.Use(middleware.TokenCheck()).PUT("/update", service.Update)
	blogGroup.Use(middleware.TokenCheck()).GET("/draft/:id", service.GetBlogDraft)

	draftGroup := group.Group("/blog-draft").Use(middleware.TokenCheck())
	draftGroup.GET(":id", service.GetDraft)
	draftGroup.GET("/list", service.ListDraft)
	draftGroup.POST("/create", service.CreateDraft)
	draftGroup.POST("/create-blog", service.CreateBlogFromDraft)
	draftGroup.PUT("/update", service.UpdateDraft)
}
