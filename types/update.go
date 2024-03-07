package types

type StatusUpdateBody struct {
	UUIDBody
	Status int `form:"status,default=0" json:"status"`
}
