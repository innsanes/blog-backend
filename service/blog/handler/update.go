package handler

import (
	"blog-backend/data/model"
	g "blog-backend/global"
	"blog-backend/service/blog/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestUpdate struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func Update(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseUint(idString, 10, strconv.IntSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	params := &RequestUpdate{}
	if err = ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = dao.Update(&model.Blog{
		Model: gorm.Model{
			ID: uint(id),
		},
		Name:    params.Name,
		Content: params.Content,
	})
	if err != nil {
		g.Log.Error("handler.blog.update error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}
