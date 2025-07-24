package model

import "gorm.io/gorm"

func init() {
	BuildList = append(BuildList, &Blog{})
	BuildList = append(BuildList, &Tag{})
}

type Blog struct {
	gorm.Model
	Name    string `gorm:"column:name;type:VARCHAR(30)"`
	Content string `gorm:"column:content;type:LONGTEXT"`
	Tags    []*Tag `gorm:"many2many:blog_tags;"`
}

type Tag struct {
	gorm.Model
	Name string `gorm:"column:name;type:VARCHAR(30)"`
}
