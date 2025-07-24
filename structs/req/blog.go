package req

type BlogCreate struct {
	Name    string   `json:"name" binding:"required,min=1,max=30"`
	Content string   `json:"content" binding:"required"`
	Tags    []string `json:"tags" binding:"omitempty,max=10,dive,min=1,max=20"`
}

type BlogUpdate struct {
	Id uint
	BlogUpdateBody
}

type BlogUpdateBody struct {
	Name    string   `json:"name" binding:"required,min=1,max=30"`
	Content string   `json:"content" binding:"required"`
	Tags    []string `json:"tags" binding:"omitempty,max=10,dive,min=1,max=20"`
}

type BlogDelete struct {
	Id uint
}

type BlogGet struct {
	Id uint
}

type BlogList struct {
	Tag       string `form:"tag" binding:"omitempty,min=1,max=20"` // 限定标签
	UseCursor bool   `form:"useCursor"`                            // 使用游标
	Page      int    `form:"page" binding:"omitempty,min=0"`       // [分页]: 第几页
	Size      int    `form:"size" binding:"required,min=1"`        // [分页]/[游标]: 每页大小
	Cursor    uint   `form:"cursor" binding:"omitempty,min=0"`     // [游标]: Blog的ID
	Forward   bool   `form:"forward"`                              // [游标]: 是向前还是向后
}
