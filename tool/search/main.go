package main

import (
	"blog-backend/core"
	"blog-backend/services/dao"
	"blog-backend/services/search"
	"blog-backend/structs/msearch"
	"fmt"
	"strconv"

	"github.com/meilisearch/meilisearch-go"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	log := core.NewLog()

	log.Info("开始创建Meilisearch索引...")

	// 创建博客索引
	log.Info("创建博客索引...")
	client := meilisearch.New("http://localhost:7700")
	err := search.CreateBlogIndex(client)
	if err != nil {
		log.Panic("创建索引失败", zap.Error(err))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		"root", "123456", "localhost", 3306, "blog", "utf8mb4")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Panic("连接数据库失败", zap.Error(err))
	}

	// 获取所有博客
	log.Info("获取所有博客...")
	blogs, err := dao.Blog.ListPage(db, 0, 100)
	if err != nil {
		log.Panic("获取博客失败", zap.Error(err))
	}

	log.Info("找到博客数量", zap.Int("count", len(blogs)))

	// 将博客转换为Meilisearch文档并插入
	successCount := 0
	for _, blog := range blogs {
		// 转换为Meilisearch格式
		searchDoc := &msearch.Blog{
			ID:      strconv.FormatUint(uint64(blog.ID), 10),
			Name:    blog.Name,
			Content: blog.Content,
		}

		// 插入到Meilisearch
		err := search.InsertBlog(client, searchDoc)
		if err != nil {
			log.Error("插入博客失败",
				zap.Uint("id", blog.ID),
				zap.String("name", blog.Name),
				zap.Error(err))
			continue
		}

		successCount++
		log.Debug("成功插入博客",
			zap.Uint("id", blog.ID),
			zap.String("name", blog.Name))
	}

	log.Info("Meilisearch索引创建完成",
		zap.Int("total", len(blogs)),
		zap.Int("success", successCount),
		zap.Int("failed", len(blogs)-successCount))
}
