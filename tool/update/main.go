package main

import (
	"blog-backend/core"
	req "blog-backend/structs/req"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/innsanes/conf"
	"go.uber.org/zap"
)

type UpdateConf struct {
	Host       string `conf:"host,default=localhost:8001,usage=server_url"`
	Protocol   string `conf:"protocol,default=http"`
	Id         uint   `conf:"id,usage=blog_id"`
	Name       string `conf:"name,usage=blog_name"`
	FilePath   string `conf:"filepath,usage=file_path"`
	Categories string `conf:"categories"`
}

//go:generate go run main.go
func main() {
	log := core.NewLog()
	config := core.NewConfig()
	_ = config.BeforeServe()
	c := &UpdateConf{}
	conf.RegisterConfWithName("s", c)
	_ = config.Serve()
	_ = config.AfterServe()
	url := fmt.Sprintf("%s://%s/blog/%d", c.Protocol, c.Host, c.Id)
	contentBytes, err := os.ReadFile(c.FilePath)
	if err != nil {
		log.Panic("读取文件失败", zap.Error(err))
	}
	payload := req.BlogUpdateBody{
		Name:       c.Name,
		Content:    string(contentBytes),
		Categories: strings.Split(c.Categories, ","),
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Panic("JSON 序列化失败", zap.Error(err))
	}

	r, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Panic("创建请求失败", zap.Error(err))
	}
	r.Header.Set("Content-Type", "application/json") // 设置请求头

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Panic("发送请求失败", zap.Error(err))
	}
	defer resp.Body.Close()

	log.Info("服务器响应状态", zap.String("status", resp.Status))
	body, _ := io.ReadAll(resp.Body)
	log.Info("服务器响应内容", zap.String("body", string(body)))
}
