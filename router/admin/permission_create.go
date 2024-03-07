package admin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type permissionCreateBody struct {
	Value string `form:"value" json:"value" binding:"required,min=6,max=64,startswith=perm_"`
	Group string `form:"group" json:"group" binding:"required,min=4,max=32"`
	Desc  string `form:"desc" json:"desc" binding:"required,min=1,max=32"`
}

func PermissionCreate(c *gin.Context) {
	var body permissionCreateBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)

	pm := &model.Perm{
		Value: body.Value,
		Group: body.Group,
		Desc:  body.Desc,
	}

	if err := db.Model(pm).Create(pm).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.Ok(c)
}
