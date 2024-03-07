package plugin

import (
	"time"

	"github.com/coorify/backend/option"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(opt *option.Option) gin.HandlerFunc {
	expire := opt.Jwt.Expire

	cfg := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "token", "signature"},
		AllowCredentials: false,
		MaxAge:           time.Duration(expire * int(time.Second)),
		AllowAllOrigins:  true,
		ExposeHeaders:    []string{"token", "signature"},
	}

	return cors.New(cfg)
}
