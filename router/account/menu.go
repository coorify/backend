package account

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/perm"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountMenu(c *gin.Context) {
	cas := c.MustGet(field.SYS_JWTCLAMIS).(*jwt.Clamis)
	pc := c.MustGet(field.SYS_PERMCACHE).(perm.PermCache)

	pmvs := pc.Get(cas.UUID, c)
	pms := make([]string, 0)
	for _, p := range pmvs {
		pms = append(pms, p.Value)
	}
	// public menu
	pms = append(pms, "")

	var mns []*model.Menu
	db := c.MustGet(field.SYS_DB).(*gorm.DB)

	if err := db.Model(&model.Menu{}).Find(&mns, "perm IN ? and status=1", pms).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.OkWithPayload(mns, c)
}
