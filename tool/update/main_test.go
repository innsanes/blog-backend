package main_test

import (
	"blog-backend/structs/req"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	_, err := os.Stat("./test.md")
	if err == nil {
		return
	}

	if os.IsNotExist(err) {
		content := []byte("# 这是一个测试 Markdown 文件\n\n由 Go 程序自动创建。\n")
		err = os.WriteFile("./test.md", content, 0644)
		if err != nil {
			return
		}
	}
}

func TestUpdate(t *testing.T) {
	payload := req.BlogUpdateBody{
		Name:    "test",
		Content: "md213123", // 将文件内容转换为字符串
		Tags:    []string{"tag2", "tag3", "tag4"},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Panic(err)
	}
	r, err := http.NewRequest("PUT", "http://localhost:8200/blog/2", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Panic(err)
	}
	r.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Printf("Status:%v Body:%v\n", resp.Status, string(body))
}
