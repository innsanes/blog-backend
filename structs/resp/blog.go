package resp

type Blog struct {
	Id         uint     `json:"id"`
	Name       string   `json:"name"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
	View       int64    `json:"view"`
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
	Categories []string `json:"categories"`
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
}
