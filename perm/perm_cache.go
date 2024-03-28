package perm

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/logger"
	"github.com/coorify/backend/model"
	"github.com/coorify/go-value"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

type PermCache interface {
	Get(uid uuid.UUID, c *gin.Context) []*PermValue
	Has(uid uuid.UUID, val *PermValue, c *gin.Context) bool
}

type permCache struct {
	grp *singleflight.Group
}

func NewPermCache() PermCache {
	return &permCache{
		grp: new(singleflight.Group),
	}
}

func (p *permCache) Get(uid uuid.UUID, c *gin.Context) []*PermValue {
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

func (p *permCache) Has(uid uuid.UUID, val *PermValue, c *gin.Context) bool {
	redisEnable := value.MustGet(c.MustGet(field.SYS_OPTION), "Redis.Enable").(bool)
	redisKey := fmt.Sprintf("PC-%s", strings.ToUpper(uid.String()))
	// redisExp := fmt.Sprintf("pce-%s", uid.String())
	if redisEnable {
		redis := c.MustGet(field.SYS_REDIS).(*redis.Client)
		has, err := redis.SMIsMember(context.Background(), redisKey, val.Value).Result()
		if err != nil {
			logger.Error(err)
			return false
		}

		if has[0] {
			return true
		}

		// _, err = redis.Get(context.Background(), redisExp).Result()
		// if err == nil {
		// 	return has[0]
		// }
	}

	pvs, err, _ := p.grp.Do(redisKey, func() (any, error) {
		pms := make([]string, 0)
		pvs := p.Get(uid, c)

		if redisEnable {
			redis := c.MustGet(field.SYS_REDIS).(*redis.Client)

			redis.Del(context.Background(), redisKey).Result()
			for _, pv := range pvs {
				redis.SAdd(context.Background(), redisKey, pv.Value).Result()
			}
			// redis.Set(context.Background(), redisExp, "", 30*time.Minute).Result()
			redis.Expire(context.Background(), redisKey, 30*time.Minute).Result()
		}

		for _, pv := range pvs {
			pms = append(pms, pv.Value)
		}
		return pms, nil
	})

	if err != nil {
		logger.Error(err)
		return false
	}

	for _, pv := range pvs.([]string) {
		if pv == val.Value {
			return true
		}
	}

	return false
}
