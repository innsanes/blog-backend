package main

import (
	"blog-backend/core"
	"blog-backend/services/dao"
	"blog-backend/structs/model"
	"fmt"
	"log"
)

func main() {
	// 初始化数据库连接
	morm := core.NewMOrm()
	if err := morm.BeforeServe(); err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}
	if err := morm.Serve(); err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	fmt.Println("开始修复博客视图记录...")

	// 获取所有博客
	blogs, err := dao.Blog.ListPage(morm.DB, 0, 1000) // 获取前1000条博客
	if err != nil {
		log.Fatalf("获取博客列表失败: %v", err)
	}

	fmt.Printf("找到 %d 个博客\n", len(blogs))

	// 为每个博客创建视图记录
	createdCount := 0
	skippedCount := 0

	for _, blog := range blogs {
		// 检查是否已经存在视图记录
		var existingView model.View
		err := morm.DB.Where("viewer_typer = ? AND viewer_id = ?", "blog", blog.ID).First(&existingView).Error
		
		if err != nil {
			// 不存在视图记录，创建新的
			err = dao.View.Create(morm.DB, "blog", blog.ID)
			if err != nil {
				log.Printf("为博客 ID %d 创建视图记录失败: %v", blog.ID, err)
				continue
			}
			createdCount++
			fmt.Printf("✓ 为博客 ID %d 创建视图记录\n", blog.ID)
		} else {
			// 已存在视图记录，跳过
			skippedCount++
			fmt.Printf("- 博客 ID %d 已有视图记录，跳过\n", blog.ID)
		}
	}

	fmt.Printf("\n修复完成！\n")
	fmt.Printf("新创建的视图记录: %d\n", createdCount)
	fmt.Printf("跳过的博客（已有视图记录）: %d\n", skippedCount)
	fmt.Printf("总计处理的博客: %d\n", len(blogs))
} 