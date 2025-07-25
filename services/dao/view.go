package dao

import (
	"blog-backend/structs/model"
	"gorm.io/gorm"
)

var View IView = &ViewDao{}

type ViewDao struct{}

type IView interface {
	Create(db *gorm.DB, viewerType string, viewerId uint) (err error)
	Increment(db *gorm.DB, viewerType string, viewerId uint) (err error)
}

func (s *ViewDao) Create(db *gorm.DB, viewerType string, viewerId uint) (err error) {
	mView := &model.View{
		ViewerID:   viewerId,
		ViewerType: viewerType,
	}
	err = db.Model(&model.View{}).Create(mView).Error
	return
}

func (s *ViewDao) Increment(db *gorm.DB, viewerType string, viewerId uint) (err error) {
	err = db.Model(&model.View{}).
		Where("viewer_typer = ?", viewerType).
		Where("viewer_id = ?", viewerId).
		UpdateColumn("count", gorm.Expr("count + ?", 1)).Error
	return
}
