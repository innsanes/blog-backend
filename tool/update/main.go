package main

import (
	"blog-backend/core"
	"blog-backend/service/blog/handler"
	"bytes"
	"encoding/json"
	"github.com/innsanes/conf"
	"io"
	"net/http"
	"os"
)

type UpdateConf struct {
	Host     string `conf:"host,default=localhost:8001,usage=server_url"`
	Protocol string `conf:"protocol,default=http"`
	Id       uint   `conf:"id,usage=blog_id"`
	Name     string `conf:"name,usage=blog_name"`
	FilePath string `conf:"filepath,usage=file_path"`
}

func main() {
	log := core.NewLog()
	config := core.NewConfig()
	_ = config.BeforeServe()
	c := &UpdateConf{}
	conf.RegisterConfWithName("s", c)
	_ = config.Serve()
	_ = config.AfterServe()
	url := c.Protocol + "://" + c.Host + "/blog/update"
	contentBytes, err := os.ReadFile(c.FilePath)
	if err != nil {
		log.Panic("读取文件失败: %v", err)
	}
	payload := handler.RequestUpdate{
		ID:      c.Id,
		Name:    c.Name,
		Content: string(contentBytes),
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Panic("JSON 序列化失败: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
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

	log.Info("服务器响应状态: %v", resp.Status)
	body, _ := io.ReadAll(resp.Body)
	log.Info("服务器响应内容: %v", string(body))
}
