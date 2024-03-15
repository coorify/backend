package perm

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/logger"
	"github.com/coorify/backend/model"
	"github.com/coorify/go-value"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type PermCache interface {
	Get(uid uuid.UUID, force bool, c *gin.Context) []*PermValue
	Has(uid uuid.UUID, val *PermValue, c *gin.Context) bool
}

type permCache struct{}

func NewPermCache() PermCache {
	return &permCache{}
}

func (p *permCache) Get(uid uuid.UUID, force bool, c *gin.Context) []*PermValue {
	return p.fromDB(uid, c)
}

func (p *permCache) Has(uid uuid.UUID, val *PermValue, c *gin.Context) bool {
	pms := p.Get(uid, false, c)
	for _, p := range pms {
		if p.Value == val.Value {
			return true
		}
	}
	return false
}

func (p *permCache) fromDB(uid uuid.UUID, c *gin.Context) []*PermValue {
	rs := make([]*PermValue, 0)
	db := c.MustGet(field.SYS_DB).(*gorm.DB)

	admUsername := value.MustGet(c.MustGet(field.SYS_OPTION), "Admin.Username").(string)

	act := &model.Account{
		UUID: uid,
	}

	if err := db.Model(act).Preload("Roles.Perms").Where(act).First(act).Error; err != nil {
		logger.Error(err)
		return rs
	}

	if len(act.Roles) == 0 && act.Username == admUsername {
		pms := AllPerm(true)
		for _, p := range pms {
			rs = append(rs, &p.PermValue)
		}

		var gps []*model.Perm
		if err := db.Model(&model.Perm{}).Distinct("value").Find(&gps).Error; err == nil {
			for _, p := range gps {
				rs = append(rs, &PermValue{Value: p.Value})
			}
		}
	} else {
		for _, rl := range act.Roles {
			if rl.Status&1 == 0 {
				continue
			}
			for _, p := range rl.Perms {
				rs = append(rs, &PermValue{Value: p.Perm})
			}
		}
	}

	return rs
}
