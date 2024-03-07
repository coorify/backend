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

type menuFindBody struct {
	types.PageBody
	types.KeywordBody
	Category string `form:"category,default=name" json:"category" binding:"oneof=perm name group"`
}

func MenuFind(c *gin.Context) {
	var body menuFindBody
	if err := c.BindQuery(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}
	var total int64
	var mns []*model.Menu
	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	db = db.Model(&model.Menu{})
	if len(body.Keyword) != 0 {
		db = db.Where(fmt.Sprintf("`%s` like ?", body.Category), fmt.Sprintf("%%%s%%", body.Keyword))
	}
	if err := db.Count(&total).Order("created_at desc").Offset((body.Page - 1) * body.Size).Limit(body.Size).Find(&mns).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.OkWithPayload(&types.PageReply{
		Page:  body.Page,
		Size:  body.Size,
		Total: int(total),
		Data:  mns,
	}, c)
}
