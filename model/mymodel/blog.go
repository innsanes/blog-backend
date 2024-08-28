package mymodel

import "gorm.io/gorm"

func init() {
	BuildList = append(BuildList, &Blog{})
}

type Blog struct {
	gorm.Model
	Name    string `gorm:"column:name;unique;type:VARCHAR(30)"`
	Content string `gorm:"column:content;type:LONGTEXT"`

	// 暂时不添加作者
	//AuthorID uint  `gorm:"column:author_id"`
	//Author   *User `gorm:"foreignKey:AuthorID;references:ID"`
}
