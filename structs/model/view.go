package model

import "gorm.io/gorm"

func init() {
	BuildList = append(BuildList, &View{})
}

type View struct {
	gorm.Model
	Count      int64  `gorm:"column:count;type:bigint"`
	ViewerID   uint   `gorm:"column:viewer_id"`
	ViewerType string `gorm:"column:viewer_typer"`
}
