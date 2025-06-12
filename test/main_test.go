package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
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
	resp, err := http.DefaultClient.Post("http://localhost:8000/blog/create", "application/json", bytes.NewReader(marshal))
	assert.Nil(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	resp, err = http.DefaultClient.Post("http://localhost:8001/blog/create", "application/json", bytes.NewReader(marshal))
	assert.Nil(t, err)
	assert.Equal(t, 201, resp.StatusCode)
}
