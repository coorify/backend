package admin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/logger"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/coorify/backend/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RolePermissions(c *gin.Context) {
	var body types.UUIDBody
	if err := c.BindQuery(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		logger.Error(err)
		return
	}

	var pms []*model.RolePerm
	db := c.MustGet(field.SYS_DB).(*gorm.DB)

	md := &model.RolePerm{
		RoleId: body.ToUUID(),
	}

	if err := db.Model(md).Where(md).Find(&pms).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	vpms := make([]string, 0)
	for _, p := range pms {
		vpms = append(vpms, p.Perm)
	}

	reply.OkWithPayload(vpms, c)
}
