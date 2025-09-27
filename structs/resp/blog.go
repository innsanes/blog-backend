package resp

type Blog struct {
	Id         uint     `json:"id"`
	Name       string   `json:"name"`
	Summary    string   `json:"summary"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
	View       int64    `json:"view"`
	Characters int64    `json:"characters"`
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
}

type BlogList struct {
	Data  []BlogListItem `json:"data"`
	Count int64          `json:"count"`
}

type BlogListItem struct {
	Id         uint     `json:"id"`
	Name       string   `json:"name"`
	Summary    string   `json:"summary"`
	Categories []string `json:"categories"`
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
}

type BlogSearchList struct {
	Data  []BlogSearchListItem `json:"data"`
	Count int64                `json:"count"`
}

type BlogSearchListItem struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Summary    string `json:"summary"`
	MatchCount int32  `json:"matchCount"`
}
