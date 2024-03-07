package admin

import (
	"fmt"
	"slices"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/perm"
	"github.com/coorify/backend/reply"
	"github.com/coorify/backend/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PermissionGroup(c *gin.Context) {
	var body types.KeywordBody
	if err := c.ShouldBindQuery(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	// 自定义权限
	var gps []*model.Perm
	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	db = db.Model(&model.Perm{}).Distinct("`group`")
	if len(body.Keyword) != 0 {
		db = db.Where("`group` like ?", fmt.Sprintf("%s%%", body.Keyword))
	}
	if err := db.Find(&gps).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	// 系统权限
	gpsn := perm.GetGroups(body.Keyword)
	// 合并所有权限
	for _, p := range gps {
		gpsn = append(gpsn, p.Group)
	}

	slices.Sort(gpsn)
	reply.OkWithPayload(slices.Compact(gpsn), c)
}
