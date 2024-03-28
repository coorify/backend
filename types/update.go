package types

import (
	"github.com/coorify/backend/field"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StatusUpdateBody struct {
	UUIDBody
	Status int  `form:"status" json:"status" binding:"required"`
	Enable bool `form:"enable,default=false" json:"enable"`
}

func StatusUpdate(c *gin.Context, model interface{}) error {
	var body StatusUpdateBody
	if err := c.ShouldBind(&body); err != nil {
		return err
	}

	expr := gorm.Expr("status | ?", body.Status)
	if !body.Enable {
		expr = gorm.Expr("status & ~?", body.Status)
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)

	return db.Model(model).Where("uuid = ?", body.ToUUID()).Update("status", expr).Error
}
