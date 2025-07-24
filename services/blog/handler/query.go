package handler

import (
	g "blog-backend/global"
	"blog-backend/services/blog/dao"
	"blog-backend/structs/model"
	"blog-backend/structs/req"
	"blog-backend/structs/tod"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	params := ctx.Param("id")
	if params == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	id, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blog, err := dao.Get(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := tod.Blog(blog)
	ctx.JSON(http.StatusOK, &response)
}

func List(ctx *gin.Context) {
	params := &req.BlogList{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var (
		blogs []*model.Blog
		count int64
		err   error
	)
	if params.Tag == "" {
		count, err = dao.Count()
	} else {
		count, err = dao.CountWithTag(params.Tag)
	}
	if err != nil {
		g.Log.Error("[Blog.ListCursorForward] error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if params.UseCursor {
		if params.Forward {
			if params.Tag == "" {
				blogs, err = dao.ListCursorForward(params.Cursor, params.Size)
			} else {
				blogs, err = dao.ListCursorForwardWithTag(params.Cursor, params.Size, params.Tag)
			}
			if err != nil {
				g.Log.Error("[Blog.ListCursorForward] error: %v", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			if params.Tag == "" {
				blogs, err = dao.ListCursorBackward(params.Cursor, params.Size)
			} else {
				blogs, err = dao.ListCursorBackwardWithTag(params.Cursor, params.Size, params.Tag)
			}
			if err != nil {
				g.Log.Error("[Blog.ListCursorBackward] error: %v", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	} else {
		if params.Tag == "" {
			blogs, err = dao.ListPage(params.Page, params.Size)
		} else {
			blogs, err = dao.ListPageWithTag(params.Page, params.Size, params.Tag)
		}
		if err != nil {
			g.Log.Error("[Blog.ListPage] error: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	response := tod.BlogList(blogs)
	response.Count = count
	ctx.JSON(http.StatusOK, &response)

}
