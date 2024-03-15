package plugin

import (
	"time"

	"github.com/coorify/go-value"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(opt interface{}) gin.HandlerFunc {
	expire := value.MustGet(opt, "Jwt.Expire").(int)

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
