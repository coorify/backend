package oauth

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/emmansun/gmsm/sm2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type profileByJwtBody struct {
	Token     string `form:"jwt" json:"jwt" binding:"required"`
	Signature string `form:"signature" json:"signature" binding:"required"`
}

func ProfileByJwt(c *gin.Context) {
	var body profileByJwtBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	bcsin, err := hex.DecodeString(body.Signature)
	if err != nil {
		reply.Fail(c)
		return
	}

	raws := []byte(body.Token)
	pub := c.MustGet(field.SYS_SIGPUBKEY).(*ecdsa.PublicKey)
	if !sm2.VerifyASN1WithSM2(pub, nil, raws, bcsin) {
		reply.Fail(c)
		return
	}

	jec := c.MustGet(field.SYS_JWT).(jwt.JwtEncoding)
	cas, ok := jec.Decode(body.Token)
	if !ok {
		reply.Fail(c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	md := &model.Account{
		UUID: cas.UUID,
	}

	if err := db.Model(md).Where(md).First(md).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.OkWithPayload(md, c)
}
