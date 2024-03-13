package admin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/coorify/backend/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type accountRoleUpdateBody struct {
	types.UUIDBody
	Values []string `form:"values" json:"values" binding:"required"`
}

func AccountRoleUpdate(c *gin.Context) {
	var body accountRoleUpdateBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	act := &model.Account{
		UUID: body.ToUUID(),
	}

	if err := db.Model(act).Where(act).First(act).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	rls := make([]*model.Role, 0)
	if err := db.Model(&model.Role{}).Find(&rls, "uuid IN ?", body.Values).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	if len(rls) > 0 {
		if err := db.Model(act).Association("Roles").Replace(rls); err != nil {
			reply.FailWithMessage(err.Error(), c)
			return
		}
	}

	reply.Ok(c)
}
