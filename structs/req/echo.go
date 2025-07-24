package req

type Echo struct {
	Message string `json:"message" binding:"required,min=1,max=30"`
}
