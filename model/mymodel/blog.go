package mymodel

import "gorm.io/gorm"

func init() {
	BuildList = append(BuildList, &Blog{})
	BuildList = append(BuildList, &BlogDraft{})
}

type Blog struct {
	gorm.Model
	Name    string `gorm:"column:name;unique;type:VARCHAR(30)"`
	Content string `gorm:"column:content;type:LONGTEXT"`

	// 暂时不添加作者
	//AuthorID uint  `gorm:"column:author_id"`
	//Author   *User `gorm:"foreignKey:AuthorID;references:ID"`
}

type BlogDraft struct {
	gorm.Model
	BlogID  uint   `gorm:"column:blog_id;index:idx_blog_id"`
	Name    string `gorm:"column:name;unique;type:VARCHAR(30)"`
	Content string `gorm:"column:content;type:LONGTEXT"`
}
