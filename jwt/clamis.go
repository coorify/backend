package jwt

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type Clamis struct {
	UUID uuid.UUID `json:"uuid"`
	// Nickname string    `json:"nickname"`
	// Status   int       `json:"status"`
}

func (c *Clamis) ToAccount(ctx *gin.Context) (*model.Account, error) {
	db := ctx.MustGet(field.SYS_DB).(*gorm.DB)
	md := &model.Account{
		UUID: c.UUID,
	}

	if err := db.Model(md).Where(md).First(md).Error; err != nil {
		return nil, err
	}

	return md, nil
}
