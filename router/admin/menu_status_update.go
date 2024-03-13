package admin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/coorify/backend/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MenuStatusUpdate(c *gin.Context) {
	var body types.StatusUpdateBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	expr := gorm.Expr("status | ?", body.Status)
	if !body.Enable {
		expr = gorm.Expr("status & ~?", body.Status)
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)

	if err := db.Model(&model.Menu{}).Where("uuid = ?", body.ToUUID()).Update("status", expr).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.Ok(c)
}
