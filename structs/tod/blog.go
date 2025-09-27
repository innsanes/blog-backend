package tod

import (
	"blog-backend/structs/model"
	"blog-backend/structs/msearch"
	"blog-backend/structs/resp"
	"blog-backend/structs/to"
	"strconv"
	"unicode/utf8"
)

func Blog(in *model.Blog) (out resp.Blog) {
	return resp.Blog{
		Id:         in.ID,
		Name:       in.Name,
		Summary:    in.Summary,
		Content:    in.Content,
		Categories: CategoryString(in.Categories),
		View:       in.View.Count,
		Characters: int64(utf8.RuneCountInString(in.Content)),
		CreateTime: in.CreatedAt.UnixMilli(),
		UpdateTime: in.UpdatedAt.UnixMilli(),
	}
}

func BlogListItem(in *model.Blog) (out resp.BlogListItem) {
	return resp.BlogListItem{
		Id:         in.ID,
		Name:       in.Name,
		Summary:    in.Summary,
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

func ImageList(in []*model.Image) (out resp.ImageList) {
	return resp.ImageList{
		Data:  to.Slice(in, ImageListItem),
		Count: 0,
	}
}

func ImageListItem(in *model.Image) (out resp.ImageListItem) {
	return resp.ImageListItem{
		Name:       in.Name,
		MD5:        in.MD5,
		CreateTime: in.CreatedAt.UnixMilli(),
	}
}

func BlogSearchList(in []*msearch.BlogSearch) (out resp.BlogSearchList) {
	return resp.BlogSearchList{
		Data:  to.Slice(in, BlogSearchListItem),
		Count: 0,
	}
}

func BlogSearchListItem(in *msearch.BlogSearch) (out resp.BlogSearchListItem) {
	id, err := strconv.ParseUint(in.ID, 10, strconv.IntSize)
	if err != nil {
		return
	}
	return resp.BlogSearchListItem{
		Id:         uint(id),
		Name:       in.Name,
		Summary:    in.Content,
		MatchCount: in.MatchCount,
	}
}
