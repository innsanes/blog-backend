package main

import (
	"blog-backend/core"
	"blog-backend/handler/blog/handler"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type CreateConf struct {
	Host     string `conf:"host,default=localhost:8001,usage=server_url"`
	Protocol string `conf:"protocol,default=http"`
	FilePath string `conf:"filepath,usage=file_path"`
}

//go:generate go run main.go -s_filepath=test.md
func main() {
	log := core.NewLog()
	config := core.NewConfig()
	_ = config.BeforeServe()
	c := &CreateConf{}
	config.RegisterConfWithName("s", c)
	_ = config.Serve()
	_ = config.AfterServe()
	url := c.Protocol + "://" + c.Host + "/blog/create"
	filePath := c.FilePath

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Panic("文件不存在: %s err:%v", filePath, err)
	}

	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Panic("读取文件失败: %v", err)
	}

	name := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	payload := handler.RequestCreate{
		Name:    name,                 // 使用 filepath.Base 获取文件名，如 "Golangmap源码解析.md"
		Content: string(contentBytes), // 将文件内容转换为字符串
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Panic("JSON 序列化失败: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Panic("创建请求失败: %v", err)
	}
	req.Header.Set("Content-Type", "application/json") // 设置请求头

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	log.Info("服务器响应状态:", resp.Status)
	body, _ := io.ReadAll(resp.Body)
	log.Info("服务器响应内容:", string(body))
}
