package admin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/coorify/backend/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountRoles(c *gin.Context) {
	var body types.UUIDBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	act := &model.Account{
		UUID: body.ToUUID(),
	}

	if err := db.Model(act).Preload("Roles.Perms").Where(act).First(act).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.OkWithPayload(act.Roles, c)
}
