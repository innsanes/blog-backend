package tom

import (
	"blog-backend/structs/model"
	"blog-backend/structs/msearch"
	"blog-backend/structs/to"
	"strconv"

	"gorm.io/gorm"
)

func BlogSearch(in *msearch.Blog) (out *model.Blog) {
	id, err := strconv.ParseUint(in.ID, 10, strconv.IntSize)
	if err != nil {
		return
	}
	return &model.Blog{
		Model:   gorm.Model{ID: uint(id)},
		Name:    in.Name,
		Content: in.Content,
	}
}

func BlogSearchList(in []*msearch.Blog) (out []*model.Blog) {
	return to.Slice(in, BlogSearch)
}
