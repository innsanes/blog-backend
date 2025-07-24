package handler

import (
	"blog-backend/global"
	"blog-backend/services/blog/dao"
	"blog-backend/structs/model"
	"blog-backend/structs/req"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	params := &req.BlogCreate{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := dao.Create(&model.Blog{
		Name:    params.Name,
		Content: params.Content,
	}, params.Tags)
	if err != nil {
		g.Log.Error("handler.blog.create error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}
