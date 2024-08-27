package service

import (
	"blog-backend/global"
	"github.com/gin-gonic/gin"
	"github.com/innsanes/serv"
	"net/http"
	"os"
	"path/filepath"
)

type Config struct {
	Path string `conf:"config,default=./image,usage=image_file_save_path"`
}

var config Config

func init() {
	serv.RegisterBeforeServe(func() error {
		global.Config.RegisterConfWithName("image", &config)
		return nil
	})
}

func Get(ctx *gin.Context) {
	// Get the image filename from the URL parameter
	imageName := ctx.Param("name")

	// Construct the file config
	filePath := filepath.Join(config.Path, imageName)

	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		global.Log.Error("Unable to open image file err:%s", err)
		ctx.String(http.StatusNotFound, "Image not found")
		return
	}
	defer file.Close()

	// Get the file's content type
	fileStat, err := file.Stat()
	if err != nil {
		global.Log.Error("Unable to get file info err:%s", err)
		ctx.String(http.StatusInternalServerError, "Unable to get file info")
		return
	}

	// Set the content type header
	ctx.Header("Content-Type", "image/"+filepath.Ext(fileStat.Name())[1:])

	// Send the file as the response
	http.ServeFile(ctx.Writer, ctx.Request, filePath)
}
