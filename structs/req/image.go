package req

import "mime/multipart"

type ImageCreate struct {
	Name string                `form:"name" binding:"required,min=1,max=32"`
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type ImageGet struct {
	Path string `form:"path"`
}

type ImageList struct {
	Page int `form:"page" binding:"omitempty,min=0" json:"page"`
	Size int `form:"size" binding:"required,min=1" json:"size"`
}
