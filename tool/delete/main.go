package main

import (
	"blog-backend/core"
	"fmt"
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
	url := fmt.Sprintf("%s://%s/blog/%d", c.Protocol, c.Host, c.Id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Panic("请求删除失败: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("请求删除失败: %v", err)
	}
	defer resp.Body.Close()

	log.Info("服务器响应状态: %v", resp.Status)
	body, _ := io.ReadAll(resp.Body)
	log.Info("服务器响应内容: %v", string(body))
}
