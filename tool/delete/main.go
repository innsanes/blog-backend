package main

import (
	"blog-backend/core"
	"blog-backend/service/blog/handler"
	"bytes"
	"encoding/json"
	"github.com/innsanes/conf"
	"io"
	"net/http"
)

type DeleteConf struct {
	Host     string `conf:"host,default=localhost:8001,usage=server_url"`
	Protocol string `conf:"protocol,default=http"`
	Id       uint   `conf:"id,usage=blog_id"`
}

func main() {
	log := core.NewLog()
	config := core.NewConfig()
	_ = config.BeforeServe()
	c := &DeleteConf{}
	conf.RegisterConfWithName("s", c)
	_ = config.Serve()
	_ = config.AfterServe()
	url := c.Protocol + "://" + c.Host + "/blog/delete"
	payload := handler.RequestDelete{
		Id: c.Id,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Panic("JSON 序列化失败: %v", err)
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Panic("请求删除失败: %v", err)
	}
	req.Header.Set("Content-Type", "application/json") // 设置请求头

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic("请求删除失败: %v", err)
	}
	defer resp.Body.Close()

	log.Info("服务器响应状态: %v", resp.Status)
	body, _ := io.ReadAll(resp.Body)
	log.Info("服务器响应内容: %v", string(body))
}
