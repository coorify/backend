package admin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type menuCreateBody struct {
	Name  string `form:"name" json:"name" binding:"required,min=4,max=16"`
	Group string `form:"group" json:"group" binding:"required,min=4,max=16"`
	Pos   string `form:"pos,default=left" json:"pos" binding:"oneof=left top"`
	Icon  string `form:"icon" json:"icon" binding:"max=32"`
	Perm  string `form:"perm" json:"perm" binding:"max=64,omitempty,startswith=perm_"`
	Link  string `form:"link" json:"link" binding:"required,min=7,max=256"`
	Auth  int    `form:"auth,default=0" json:"auth"`
}

func MenuCreate(c *gin.Context) {
	var body menuCreateBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	mn := &model.Menu{
		Name:  body.Name,
		Group: body.Group,
		Icon:  body.Icon,
		Perm:  body.Perm,
		Link:  body.Link,
		Auth:  body.Auth,
		Pos:   body.Pos,
	}

	if err := db.Model(mn).Create(mn).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.Ok(c)
}
