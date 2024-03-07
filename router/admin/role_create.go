package admin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type roleCreateBody struct {
	Name string `form:"name" json:"name" binding:"required,min=4,max=16"`
	Desc string `form:"desc" json:"desc" binding:"required,min=1,max=256"`
}

func RoleCreate(c *gin.Context) {
	var body roleCreateBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	rl := &model.Role{
		Name: body.Name,
		Desc: body.Desc,
	}

	if err := db.Model(rl).Create(rl).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.Ok(c)
}
