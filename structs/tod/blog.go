package tod

import (
	"blog-backend/structs/model"
	"blog-backend/structs/resp"
	"blog-backend/structs/to"
)

func Blog(in *model.Blog) (out resp.Blog) {
	return resp.Blog{
		Id:         in.ID,
		Name:       in.Name,
		Content:    in.Content,
		Categories: CategoryString(in.Categories),
		View:       in.View.Count,
		CreateTime: in.CreatedAt.UnixMilli(),
		UpdateTime: in.UpdatedAt.UnixMilli(),
	}
}

func BlogListItem(in *model.Blog) (out resp.BlogListItem) {
	return resp.BlogListItem{
		Id:         in.ID,
		Name:       in.Name,
		Categories: CategoryString(in.Categories),
		CreateTime: in.CreatedAt.UnixMilli(),
		UpdateTime: in.UpdatedAt.UnixMilli(),
	}
}

func BlogList(in []*model.Blog) (out resp.BlogList) {
	return resp.BlogList{
		Data:  to.Slice(in, BlogListItem),
		Count: 0,
	}
}

func CategoryString(in []*model.Category) []string {
	return to.Slice(in, func(elem *model.Category) string { return elem.Name })
}
