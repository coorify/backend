package types

type PageBody struct {
	Page int `form:"page,default=1" json:"page" binding:"required"`
	Size int `form:"size,default=20" json:"size" binding:"required"`
}

type PageReply struct {
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
