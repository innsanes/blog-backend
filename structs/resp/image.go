package resp

type ImageCreate struct {
	MD5 string `json:"md5"`
}

type ImageList struct {
	Data  []ImageListItem `json:"data"`
	Count int64           `json:"count"`
}

type ImageListItem struct {
	Name       string `json:"name"`
	MD5        string `json:"md5"`
	CreateTime int64  `json:"create_time"`
}
