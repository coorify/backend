package admin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/coorify/backend/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountStatusUpdate(c *gin.Context) {
	var body types.StatusUpdateBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	md := &model.Account{
		UUID: body.ToUUID(),
	}

	if err := db.Model(md).Where(md).Update("status", body.Status).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
	}

	reply.Ok(c)
}
