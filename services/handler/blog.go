package handler

import (
	g "blog-backend/global"
	"blog-backend/services/service"
	"blog-backend/structs/req"
	"blog-backend/structs/tod"
	"net/http"
	"strconv"

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
		g.Log.Error("%s", err.Error())
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
		g.Log.Error("handler.blog.update error: %v", err)
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
	blog, err := service.Blog.Get(request)
	if err != nil {
		g.Log.Error("handler.blog.get error: %v", err)
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
	blogs, err := service.Blog.List(params)
	if err != nil {
		g.Log.Error("handler.blog.list error: %v", err)
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
		g.Log.Error("handler.blog.delete error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
