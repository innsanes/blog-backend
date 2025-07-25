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
		Tags:       TagString(in.Tags),
		View:       in.View.Count,
		CreateTime: in.CreatedAt.UnixMilli(),
		UpdateTime: in.UpdatedAt.UnixMilli(),
	}
}

func BlogListItem(in *model.Blog) (out resp.BlogListItem) {
	return resp.BlogListItem{
		Id:         in.ID,
		Name:       in.Name,
		Tags:       TagString(in.Tags),
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

func TagString(in []*model.Tag) []string {
	return to.Slice(in, func(elem *model.Tag) string { return elem.Name })
}
