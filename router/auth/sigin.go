package auth

import (
	"encoding/base64"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type siginBody struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type siginReply struct {
	Token  string `json:"token"`
	Expire int    `json:"expire"`
}

func Sigin(c *gin.Context) {
	var body siginBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	jt := c.MustGet(field.SYS_JWT).(jwt.JwtEncoding)

	cd := &model.Account{
		Username: body.Username,
	}

	rs := &model.Account{}
	if err := db.Model(cd).Where(cd).First(rs).Error; err != nil {
		reply.FailWithMessage("账号或密码错误", c)
		return
	}

	bytes, err := base64.StdEncoding.DecodeString(body.Password)
	if err != nil || !rs.Verify(string(bytes)) {
		reply.FailWithMessage("账号或密码错误", c)
		return
	}

	if rs.Status&1 == 0 {
		reply.FailWithMessage("账号已被冻结", c)
		return
	}

	token := jt.MustEncode(&jwt.Clamis{
		UUID: rs.UUID,
		// Nickname: rs.Nickname,
		// Status:   rs.Status,
	})
	expire := jt.Expire()

	reply.OkWithPayload(&siginReply{
		Token:  token,
		Expire: expire,
	}, c)
}
