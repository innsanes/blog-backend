package handler

import (
	"blog-backend/data/model"
	"blog-backend/global"
	"blog-backend/service/blog/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestCreate struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func Create(ctx *gin.Context) {
	params := &RequestCreate{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := dao.Create(&model.Blog{
		Name:    params.Name,
		Content: params.Content,
	})
	if err != nil {
		g.Log.Error("handler.blog.create error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}
