package main_test

import (
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
