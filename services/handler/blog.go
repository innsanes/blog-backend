package handler

import (
	"blog-backend/services/service"
	"blog-backend/structs/req"
	"blog-backend/structs/tod"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func BlogCreate(ctx *gin.Context) {
	params := &req.BlogCreate{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.Blog.Create(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}

func BlogUpdate(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseUint(idString, 10, strconv.IntSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	params := &req.BlogUpdateBody{}
	if err = ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request := &req.BlogUpdate{
		Id:             uint(id),
		BlogUpdateBody: *params,
	}
	err = service.Blog.Update(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

func BlogGet(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseUint(idString, 10, strconv.IntSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request := &req.BlogGet{
		Id: uint(id),
	}
	ctx.Set("request", fmt.Sprintf("%s=%s", "id", idString))
	blog, err := service.Blog.Get(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := tod.Blog(blog)
	ctx.JSON(http.StatusOK, response)
}

func BlogGetAdmin(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseUint(idString, 10, strconv.IntSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request := &req.BlogGet{
		Id: uint(id),
	}
	blog, err := service.Blog.GetAdmin(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := tod.Blog(blog)
	ctx.JSON(http.StatusOK, response)
}

func BlogList(ctx *gin.Context) {
	params := &req.BlogList{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	kvs := []string{
		fmt.Sprintf("%s=%s", "category", params.Category),
		fmt.Sprintf("%s=%v", "userCursor", params.UseCursor),
		fmt.Sprintf("%s=%d", "page", params.Page),
		fmt.Sprintf("%s=%d", "size", params.Size),
		fmt.Sprintf("%s=%d", "cursor", params.Cursor),
		fmt.Sprintf("%s=%v", "forward", params.Forward),
	}
	ctx.Set("request", strings.Join(kvs, ","))
	blogs, err := service.Blog.List(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := tod.BlogList(blogs)
	ctx.JSON(http.StatusOK, response)
}

func BlogDelete(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseUint(idString, 10, strconv.IntSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request := &req.BlogDelete{
		Id: uint(id),
	}
	err = service.Blog.Delete(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func BlogSearch(ctx *gin.Context) {
	params := &req.BlogSearch{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Set("request", fmt.Sprintf("%s=%s", "search", params.Search))
	blogs, err := service.Blog.Search(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := tod.BlogSearchList(blogs)
	ctx.JSON(http.StatusOK, response)
}
