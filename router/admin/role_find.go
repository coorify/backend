package admin

import (
	"fmt"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/coorify/backend/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoleFind(c *gin.Context) {
	var body types.FindBody
	if err := c.BindQuery(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	var total int64
	var rls []*model.Role
	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	db = db.Model(&model.Role{}).Preload("Perms")
	if len(body.Keyword) != 0 {
		db = db.Where("`name` like ?", fmt.Sprintf("%%%s%%", body.Keyword))
	}

	if err := db.Count(&total).Order("created_at desc").Offset((body.Page - 1) * body.Size).Limit(body.Size).Find(&rls).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.OkWithPayload(&types.PageReply{
		Page:  body.Page,
		Size:  body.Size,
		Total: int(total),
		Data:  rls,
	}, c)
}
