package plugin

import (
	"context"
	"fmt"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/logger"
	"github.com/coorify/go-value"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Redis(opt interface{}) gin.HandlerFunc {
	enable := value.MustGet(opt, "Redis.Enable").(bool)
	host := value.MustGet(opt, "Redis.Host").(string)
	port := value.MustGet(opt, "Redis.Port").(int)
	pswd := value.MustGet(opt, "Redis.Password").(string)
	db := value.MustGet(opt, "Redis.DB").(int)

	if !enable {
		return func(c *gin.Context) {}
	}

	ins := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: pswd,
		DB:       db,
	})

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
