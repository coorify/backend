package plugin

import (
	"time"

	"github.com/coorify/backend/option"
	"github.com/coorify/go-value"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(opt interface{}) gin.HandlerFunc {
	o := value.MustGet(opt, "Jwt").(option.JwtOption)
	expire := o.Expire

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
