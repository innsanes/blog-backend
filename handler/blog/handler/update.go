package handler

import (
	"blog-backend/data/model"
	g "blog-backend/global"
	"blog-backend/handler/blog/dao"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestUpdate struct {
	ID      uint   `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func Update(ctx *gin.Context) {
	params := &RequestUpdate{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := dao.Update(&model.Blog{
		Model: gorm.Model{
			ID: params.ID,
		},
		Name:    params.Name,
		Content: params.Content,
	})
	if err != nil {
		g.Log.Error("handler.blog.update error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}
