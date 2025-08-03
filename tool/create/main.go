package main

import (
	"blog-backend/core"
	"blog-backend/structs/req"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"fmt"

	"github.com/innsanes/conf"
	"go.uber.org/zap"
)

type CreateConf struct {
	Host       string `conf:"host,default=localhost:8001,usage=server_url"`
	Protocol   string `conf:"protocol,default=http"`
	Name       string `conf:"name,default=test,usage=blog_name"`
	FilePath   string `conf:"filepath,default=test.md,usage=file_path"`
	Categories string `conf:"categories,default=,usage=tags"`
}

func main() {
	log := core.NewLog()
	config := core.NewConfig()
	_ = config.BeforeServe()
	c := &CreateConf{}
	conf.RegisterConfWithName("s", c)
	_ = config.Serve()
	_ = config.AfterServe()
	url := fmt.Sprintf("%s://%s/blog", c.Protocol, c.Host)
	filePath := c.FilePath

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Panic("文件不存在", zap.String("filePath", filePath), zap.Error(err))
	}

	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Panic("读取文件失败", zap.Error(err))
	}

	payload := req.BlogCreate{
		Name:       c.Name,
		Content:    string(contentBytes), // 将文件内容转换为字符串
		Categories: strings.Split(c.Categories, ","),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Panic("JSON 序列化失败", zap.Error(err))
	}

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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
