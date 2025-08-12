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
	GetAdmin(in *req.BlogGet) (out *model.Blog, err error)
	List(in *req.BlogList) (out []*model.Blog, err error)
	Delete(in *req.BlogDelete) (err error)
}

func (s *BlogService) Create(in *req.BlogCreate) (err error) {
	err = g.MySQL.Transaction(func(tx *gorm.DB) (txErr error) {
		mCategories, txErr := s.findOrCreateCategories(tx, in.Categories)
		if txErr = errc.Handle("[Blog.Create] findOrCreateCategories", txErr); txErr != nil {
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
		txErr = dao.Blog.UpdateCategories(tx, m, mCategories)
		if txErr = errc.Handle("[Blog.Create] UpdateCategories", txErr); txErr != nil {
			return
		}
		txErr = dao.View.Create(tx, "blog", m.ID)
		if txErr = errc.Handle("[Blog.Create] View", txErr); txErr != nil {
			return
		}
		return
	})
	if err = errc.Handle("[Blog.Create] Transaction", err); err != nil {
		return
	}
	return
}

func (s *BlogService) findOrCreateCategories(db *gorm.DB, categories []string) (mCategories []*model.Category, err error) {
	mCategories, err = dao.Category.ListByName(db, categories)
	if err = errc.Handle("[Blog.findOrCreateCategories] ListByName", err); err != nil {
		return
	}
	if len(categories)-len(mCategories) <= 0 {
		return
	}
	cCategories := make([]string, 0, len(categories)-len(mCategories))
	for _, category := range categories {
		isContain := slices.ContainsFunc(mCategories, func(mCategory *model.Category) bool {
			return mCategory.Name == category
		})
		if !isContain {
			cCategories = append(cCategories, category)
		}
	}
	appendCategories, err := dao.Category.CreateMulti(db, cCategories)
	if err = errc.Handle("[Blog.findOrCreateCategories] CreateMulti", err); err != nil {
		return
	}
	mCategories = append(mCategories, appendCategories...)
	return
}

func (s *BlogService) Update(in *req.BlogUpdate) (err error) {
	err = g.MySQL.Transaction(func(tx *gorm.DB) (txError error) {
		mCategories, txError := s.findOrCreateCategories(tx, in.Categories)
		if txError = errc.Handle("[Blog.Update] findOrCreateCategories", txError); txError != nil {
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
		txError = dao.Blog.UpdateCategories(tx, mBlog, mCategories)
		if txError = errc.Handle("[Blog.Update] UpdateCategories", txError); txError != nil {
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
	err = dao.View.Increment(g.MySQL.DB, "blog", in.Id)
	if err = errc.Handle("[Blog.Get] Increment", err); err != nil {
		return nil, err
	}
	mBlog.View.Count += 1
	return mBlog, nil
}

func (s *BlogService) GetAdmin(in *req.BlogGet) (out *model.Blog, err error) {
	out, err = dao.Blog.Get(g.MySQL.DB, in.Id)
	if err = errc.Handle("[Blog.GetAdmin] Get", err); err != nil {
		return
	}
	return
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
	if in.Category == "" {
		out, err = dao.Blog.ListPage(db, in.Page, in.Size)
		if err = errc.Handle("[Blog.ListWithPage] ListPage", err); err != nil {
			return
		}
	} else {
		var category *model.Category
		category, err = dao.Category.GetByName(db, in.Category)
		if err = errc.Handle("[Blog.ListWithPage] GetByName", err); err != nil {
			return
		}
		out, err = dao.Blog.ListPageWithCategory(db, in.Page, in.Size, category.ID)
		if err = errc.Handle("[Blog.ListWithPage] ListPageWithCategory", err); err != nil {
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
	if in.Category == "" {
		out, err = dao.Blog.ListCursorForward(db, in.Cursor, in.Size)
		if err = errc.Handle("[Blog.ListWithCursorForward] ListCursorForward", err); err != nil {
			return
		}
	} else {
		var category *model.Category
		category, err = dao.Category.GetByName(db, in.Category)
		if err = errc.Handle("[Blog.ListWithCursorForward] GetByName", err); err != nil {
			return
		}
		out, err = dao.Blog.ListCursorForwardWithCategory(db, in.Cursor, in.Size, category.ID)
		if err = errc.Handle("[Blog.ListWithCursorForward] ListCursorForwardWithCategory", err); err != nil {
			return
		}
	}
	return
}

func (s *BlogService) ListWithCursorBackward(in *req.BlogList) (out []*model.Blog, err error) {
	db := g.MySQL.DB
	if in.Category == "" {
		out, err = dao.Blog.ListCursorBackward(db, in.Cursor, in.Size)
		if err = errc.Handle("[Blog.ListWithCursorBackward] ListCursorBackward", err); err != nil {
			return
		}
	} else {
		var category *model.Category
		category, err = dao.Category.GetByName(db, in.Category)
		if err = errc.Handle("[Blog.ListWithCursorBackward] GetByName", err); err != nil {
			return
		}
		out, err = dao.Blog.ListCursorBackwardWithCategory(db, in.Cursor, in.Size, category.ID)
		if err = errc.Handle("[Blog.ListWithCursorBackward] ListCursorBackwardWithCategory", err); err != nil {
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
