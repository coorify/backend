package types

type StatusUpdateBody struct {
	UUIDBody
	Status int  `form:"status" json:"status" binding:"required"`
	Enable bool `form:"enable,default=false" json:"enable"`
}
