package admin

import (
	"fmt"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/coorify/backend/types"
	"github.com/coorify/go-value"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type accountFindBody struct {
	types.PageBody
	types.KeywordBody
	Category string `form:"category,default=nickname" json:"category" binding:"oneof=username nickname phone"`
}

func AccountFind(c *gin.Context) {
	var body accountFindBody
	if err := c.ShouldBind(&body); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	var total int64
	var acts []*model.Account
	db := c.MustGet(field.SYS_DB).(*gorm.DB)
	admUsername := value.MustGet(c.MustGet(field.SYS_OPTION), "Admin.Username").(string)

	db = db.Model(&model.Account{}).Preload("Roles").Where("username <> ?", admUsername) // 不查询admin账号
	if len(body.Keyword) != 0 {
		db = db.Where(fmt.Sprintf("`%s` like ?", body.Category), fmt.Sprintf("%%%s%%", body.Keyword))
	}

	if err := db.Count(&total).Order("created_at desc").Offset((body.Page - 1) * body.Size).Limit(body.Size).Find(&acts).Error; err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.OkWithPayload(&types.PageReply{
		Page:  body.Page,
		Size:  body.Size,
		Total: int(total),
		Data:  acts,
	}, c)
}
