package account

import (
	"encoding/base64"

	"github.com/coorify/backend/crypto"
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type passwordUpdateBody struct {
	Old string `form:"old" json:"old" binding:"required,base64"`
	New string `form:"new" json:"new" binding:"required,base64"`
}

func PasswordUpdate(c *gin.Context) {
	var body passwordUpdateBody
	var raws []byte
	var err error

	if err = c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	raws, err = base64.StdEncoding.DecodeString(body.Old)
	if err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}
	body.Old = string(raws)

	raws, err = base64.RawStdEncoding.DecodeString(body.New)
	if err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}
	body.New = string(raws)

	if len(body.New) < 6 {
		reply.FailWithMessage("新密码长度不足", c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	cas := c.MustGet(field.SYS_JWTCLAMIS).(*jwt.Clamis)

	act, err := cas.ToAccount(c)
	if err != nil || act.Status&1 == 0 {
		reply.Fail(c)
		return
	}

	if !act.Verify(body.Old) {
		reply.FailWithMessage("旧密码不匹配", c)
		return
	}

	cd := &model.Account{
		UUID: act.UUID,
	}

	up := &model.Account{
		Password: crypto.EncodePassword(act.Username, body.New),
	}

	err = db.Model(cd).Where(cd).Updates(up).Error
	if err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.Ok(c)
}
