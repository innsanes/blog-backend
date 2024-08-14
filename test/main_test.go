package test

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	// 模拟客户端 发送http请求
	resp, err := http.DefaultClient.Get("http://localhost:8000/ping")
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
