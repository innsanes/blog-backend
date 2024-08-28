package service

import (
	"blog-backend/handler/blog/dao"
	"blog-backend/model/mymodel"
	"github.com/gin-gonic/gin"
	"net/http"
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
	err := dao.Create(&mymodel.Blog{
		Name:    params.Name,
		Content: params.Content,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
	return
}
