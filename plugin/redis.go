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

func Redis(s Server) gin.HandlerFunc {
	o := s.Option()

	enable := value.MustGet(o, "Redis.Enable").(bool)
	if !enable {
		return func(c *gin.Context) {}
	}

	opt := &redis.Options{}
	if err := merge.Merge(opt, value.MustGet(o, "Redis")); err != nil {
		panic(err)
	}

	_redis := redis.NewClient(opt)
	res, err := _redis.ClientInfo(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	logger.Infof("redis remote: %s", _redis.Options().Addr)
	logger.Infof("redis lib: %s", res.LibName)
	logger.Infof("redis ver: %s", res.LibVer)

	s.Set(field.SYS_REDIS, _redis)
	return func(c *gin.Context) {
		c.Set(field.SYS_REDIS, _redis)
	}
}
