package admin

import (
	"encoding/base64"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type accountCreateBody struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=16"`
	Password string `form:"password" json:"password" binding:"required,base64"`
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=2,max=16"`
	Phone    string `form:"phone" json:"phone" binding:"required,len=11"`
}

func AccountCreate(c *gin.Context) {
	var body accountCreateBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	raws, err := base64.StdEncoding.DecodeString(body.Password)
	if err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}
	body.Password = string(raws)

	if len(body.Password) < 6 {
		reply.FailWithMessage("新密码长度不足", c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	var rls *[]model.Role
	if err := db.Model(&model.Role{}).Where("status&2 != 0").Find(&rls).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	act := &model.Account{
		Username: body.Username,
		Nickname: body.Nickname,
		Phone:    body.Phone,
		Password: body.Password,
		Roles:    *rls,
	}

	if err := db.Model(act).Create(act).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.Ok(c)
}
