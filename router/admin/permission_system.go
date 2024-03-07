package admin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/perm"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PermissionSystem(c *gin.Context) {
	var cpms []*model.Perm
	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	if err := db.Model(&model.Perm{}).Find(&cpms).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	pms := make([]*perm.Perm, 0)
	pms = append(pms, perm.AllPerm(false)...)

	for _, p := range cpms {
		pms = append(pms, &perm.Perm{
			PermValue: perm.PermValue{
				Value: p.Value,
			},
			Group: p.Group,
			Desc:  p.Desc,
		})
	}

	reply.OkWithPayload(pms, c)
}
