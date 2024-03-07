package account

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Profile(c *gin.Context) {
	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	cas := c.MustGet(field.SYS_JWTCLAMIS).(*jwt.Clamis)

	md := &model.Account{
		UUID: cas.UUID,
	}

	if err := db.Model(md).Where(md).First(md).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.OkWithPayload(md, c)
}
