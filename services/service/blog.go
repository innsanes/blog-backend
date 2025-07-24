package service

import (
	g "blog-backend/global"
	"blog-backend/services/dao"
	"blog-backend/structs/errc"
	"blog-backend/structs/model"
	"blog-backend/structs/req"
	"slices"

	"gorm.io/gorm"
)

var Blog IBlog = &BlogService{}

type BlogService struct{}

type IBlog interface {
	Create(in *req.BlogCreate) (err error)
	Update(in *req.BlogUpdate) (err error)
	Get(in *req.BlogGet) (out *model.Blog, err error)
	List(in *req.BlogList) (out []*model.Blog, err error)
	Delete(in *req.BlogDelete) (err error)
}

func (s *BlogService) Create(in *req.BlogCreate) (err error) {
	err = g.MySQL.Transaction(func(tx *gorm.DB) (txErr error) {
		mTags, txErr := s.findOrCreateTags(tx, in.Tags)
		if txErr = errc.Handle("[Blog.Create] findOrCreateTags", txErr); txErr != nil {
			return
		}
		m := &model.Blog{
			Name:    in.Name,
			Content: in.Content,
		}
		txErr = dao.Blog.Create(tx, m)
		if txErr = errc.Handle("[Blog.Create] Create", txErr); txErr != nil {
			return
		}
		txErr = dao.Blog.UpdateTags(tx, m, mTags)
		if txErr = errc.Handle("[Blog.Create] UpdateTags", txErr); txErr != nil {
			return
		}
		return
	})
	if err = errc.Handle("[Blog.Create] Transaction", err); err != nil {
		return
	}
	return
}

func (s *BlogService) findOrCreateTags(db *gorm.DB, tags []string) (mTags []*model.Tag, err error) {
	mTags, err = dao.Tag.ListByName(db, tags)
	if err = errc.Handle("[Blog.findOrCreateTags] ListByName", err); err != nil {
		return
	}
	if len(tags)-len(mTags) <= 0 {
		return
	}
	cTags := make([]string, 0, len(tags)-len(mTags))
	for _, tag := range tags {
		isContain := slices.ContainsFunc(mTags, func(mTag *model.Tag) bool {
			return mTag.Name == tag
		})
		if !isContain {
			cTags = append(cTags, tag)
		}
	}
	appendTags, err := dao.Tag.CreateMulti(db, cTags)
	if err = errc.Handle("[Blog.findOrCreateTags] CreateMulti", err); err != nil {
		return
	}
	mTags = append(mTags, appendTags...)
	return
}

func (s *BlogService) Update(in *req.BlogUpdate) (err error) {
	err = g.MySQL.Transaction(func(tx *gorm.DB) (txError error) {
		mTags, txError := s.findOrCreateTags(tx, in.Tags)
		if txError = errc.Handle("[Blog.Update] findOrCreateTags", txError); txError != nil {
			return
		}
		mBlog := &model.Blog{
			Model: gorm.Model{
				ID: in.Id,
			},
			Name:    in.Name,
			Content: in.Content,
		}
		txError = dao.Blog.Update(tx, mBlog)
		if txError = errc.Handle("[Blog.Update] Update", txError); txError != nil {
			return
		}
		txError = dao.Blog.UpdateTags(tx, mBlog, mTags)
		if txError = errc.Handle("[Blog.Update] UpdateTags", txError); txError != nil {
			return
		}
		return
	})
	if err = errc.Handle("[Blog.Update] Transaction", err); err != nil {
		return
	}
	return
}

func (s *BlogService) Get(in *req.BlogGet) (out *model.Blog, err error) {
	mBlog, err := dao.Blog.Get(g.MySQL.DB, in.Id)
	if err = errc.Handle("[Blog.Get] Get", err); err != nil {
		return
	}
	return mBlog, nil
}

func (s *BlogService) List(in *req.BlogList) (out []*model.Blog, err error) {
	if in.UseCursor {
		out, err = s.ListWithCursor(in)
	} else {
		out, err = s.ListWithPage(in)
	}
	return
}

func (s *BlogService) ListWithPage(in *req.BlogList) (out []*model.Blog, err error) {
	db := g.MySQL.DB
	if in.Tag == "" {
		out, err = dao.Blog.ListPage(db, in.Page, in.Size)
		if err = errc.Handle("[Blog.ListWithPage] ListPage", err); err != nil {
			return
		}
	} else {
		var tag *model.Tag
		tag, err = dao.Tag.GetByName(db, in.Tag)
		if err = errc.Handle("[Blog.ListWithPage] GetByName", err); err != nil {
			return
		}
		out, err = dao.Blog.ListPageWithTag(db, in.Page, in.Size, tag.ID)
		if err = errc.Handle("[Blog.ListWithPage] ListPageWithTag", err); err != nil {
			return
		}
	}
	return
}

func (s *BlogService) ListWithCursor(in *req.BlogList) (out []*model.Blog, err error) {
	if in.Forward {
		out, err = s.ListWithCursorForward(in)
	} else {
		out, err = s.ListWithCursorBackward(in)
	}
	return
}

func (s *BlogService) ListWithCursorForward(in *req.BlogList) (out []*model.Blog, err error) {
	db := g.MySQL.DB
	if in.Tag == "" {
		out, err = dao.Blog.ListCursorForward(db, in.Cursor, in.Size)
		if err = errc.Handle("[Blog.ListWithCursorForward] ListCursorForward", err); err != nil {
			return
		}
	} else {
		var tag *model.Tag
		tag, err = dao.Tag.GetByName(db, in.Tag)
		if err = errc.Handle("[Blog.ListWithCursorForward] GetByName", err); err != nil {
			return
		}
		out, err = dao.Blog.ListCursorForwardWithTag(db, in.Cursor, in.Size, tag.ID)
		if err = errc.Handle("[Blog.ListWithCursorForward] ListCursorForwardWithTag", err); err != nil {
			return
		}
	}
	return
}

func (s *BlogService) ListWithCursorBackward(in *req.BlogList) (out []*model.Blog, err error) {
	db := g.MySQL.DB
	if in.Tag == "" {
		out, err = dao.Blog.ListCursorBackward(db, in.Cursor, in.Size)
		if err = errc.Handle("[Blog.ListWithCursorBackward] ListCursorBackward", err); err != nil {
			return
		}
	} else {
		var tag *model.Tag
		tag, err = dao.Tag.GetByName(db, in.Tag)
		if err = errc.Handle("[Blog.ListWithCursorBackward] GetByName", err); err != nil {
			return
		}
		out, err = dao.Blog.ListCursorBackwardWithTag(db, in.Cursor, in.Size, tag.ID)
		if err = errc.Handle("[Blog.ListWithCursorBackward] ListCursorBackwardWithTag", err); err != nil {
			return
		}
	}
	return
}

func (s *BlogService) Delete(in *req.BlogDelete) (err error) {
	err = dao.Blog.Delete(g.MySQL.DB, in.Id)
	if err = errc.Handle("[Blog.Delete] Delete", err); err != nil {
		return
	}
	return
}
