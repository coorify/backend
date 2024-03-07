package admin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/coorify/backend/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type rolePermissionUpdateBody struct {
	types.UUIDBody
	Values []string `form:"values" json:"values" binding:"required"`
}

func RolePermissionUpdate(c *gin.Context) {
	var body rolePermissionUpdateBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	dl := &model.RolePerm{
		RoleId: body.ToUUID(),
	}
	if err := db.Model(dl).Unscoped().Where(dl).Delete(dl).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	pms := make([]*model.RolePerm, 0)
	for _, p := range body.Values {
		pms = append(pms, &model.RolePerm{
			RoleId: body.ToUUID(),
			Perm:   p,
		})
	}

	if err := db.Model(&model.RolePerm{}).Create(pms).Error; err != nil {
		db.Model(dl).Unscoped().Where(dl).Delete(dl)
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.Ok(c)
}
