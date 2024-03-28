package plugin

import (
	"context"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/logger"
	"github.com/coorify/go-merge"
	"github.com/coorify/go-value"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Redis(o interface{}) gin.HandlerFunc {
	enable := value.MustGet(o, "Redis.Enable").(bool)

	if !enable {
		return func(c *gin.Context) {}
	}

	opt := &redis.Options{}
	if err := merge.Merge(opt, value.MustGet(o, "Redis")); err != nil {
		panic(err)
	}

	ins := redis.NewClient(opt)
	res, err := ins.ClientInfo(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	logger.Infof("redis remote: %s", ins.Options().Addr)
	logger.Infof("redis lib: %s", res.LibName)
	logger.Infof("redis ver: %s", res.LibVer)
	return func(c *gin.Context) {
		c.Set(field.SYS_REDIS, ins)
	}
}
