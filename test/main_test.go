package test

import (
	"blog-backend/handler/user/service"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

type RequestCreate struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func TestClient(t *testing.T) {
	// 模拟客户端 发送http请求
	marshal, err := json.Marshal(&RequestCreate{
		Name:    "test233",
		Content: "test3232",
	})
	assert.Nil(t, err)
	http.DefaultClient.Post("http://localhost:8000/login", "application/json", bytes.NewBuffer(marshal))
	http.DefaultClient.Post("http://localhost:8000/login", "application/json", bytes.NewBuffer(marshal))
	http.DefaultClient.Post("http://localhost:8000/login", "application/json", bytes.NewBuffer(marshal))
	http.DefaultClient.Post("http://localhost:8000/login", "application/json", bytes.NewBuffer(marshal))
	resp, err := http.DefaultClient.Post("http://localhost:8000/blog/create", "application/json", bytes.NewReader(marshal))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// 读取响应的body
	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	defer resp.Body.Close()

	// 将body内容转换为字符串
	bodyStr := string(body)

	// 验证body内容是否正确
	expectedBody := "{\"message\":\"pong\"}"
	assert.Equal(t, expectedBody, bodyStr)
}

func TestRegister(t *testing.T) {
	// 模拟客户端 发送http请求
	marshal, err := json.Marshal(&service.RequestRegister{
		Name:     "admin1",
		Password: "1234561",
	})
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Post("http://localhost:8000/register", "application/json", bytes.NewReader(marshal))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
