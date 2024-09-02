package service

import (
	"blog-backend/global"
	"blog-backend/handler/blog/dao"
	"blog-backend/model/mymodel"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

type RequestCreateDraft struct {
	BlogID  uint   `json:"blog_id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreateDraft(ctx *gin.Context) {
	params := &RequestCreateDraft{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := dao.CreateDraft(&mymodel.BlogDraft{
		BlogID:  params.BlogID,
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

type RequestCreateBlogFromDraft struct {
	ID      uint   `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreateBlogFromDraft(ctx *gin.Context) {
	params := &RequestCreateBlogFromDraft{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := global.MySQL.Transaction(func(tx *gorm.DB) error {
		blog := &mymodel.Blog{
			Name:    params.Name,
			Content: params.Content,
		}
		err := dao.Create(blog)
		if err != nil {
			return err
		}
		err = dao.UpdateDraft(&mymodel.BlogDraft{
			Model: gorm.Model{
				ID: params.ID,
			},
			BlogID:  blog.ID,
			Name:    params.Name,
			Content: params.Content,
		})
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
	return
}
