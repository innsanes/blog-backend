package msearch

type Blog struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type BlogSearch struct {
	Blog
	MatchCount int32 `json:"matchCount"`
}

type MatchPosition struct {
	Start  int `json:"start"`
	Length int `json:"length"`
}

type MatchPositions []MatchPosition

type BlogSearchMatchPositions struct {
	Name    MatchPositions `json:"name"`
	Content MatchPositions `json:"content"`
}
