package admin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type permissionDeleteBody struct {
	Value string `form:"value" json:"value" binding:"required,min=6,max=64,startswith=perm_"`
}

func PermissionDelete(c *gin.Context) {
	var body permissionDeleteBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	md := &model.Perm{
		Value: body.Value,
	}

	if err := db.Model(md).Where(md).Delete(md).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	rpm := &model.RolePerm{
		Perm: body.Value,
	}
	if err := db.Model(rpm).Unscoped().Where(rpm).Delete(rpm).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.Ok(c)
}
