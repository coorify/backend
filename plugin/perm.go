package plugin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/perm"
	"github.com/gin-gonic/gin"
)

func Perm(opt interface{}) gin.HandlerFunc {
	pc := perm.NewPermCache()

	return func(c *gin.Context) {
		c.Set(field.SYS_PERMCACHE, pc)
	}
}
