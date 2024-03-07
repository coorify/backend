package plugin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/option"
	"github.com/coorify/backend/perm"
	"github.com/gin-gonic/gin"
)

func Perm(opt *option.Option) gin.HandlerFunc {
	pc := perm.NewPermCache()

	return func(c *gin.Context) {
		c.Set(field.SYS_PERMCACHE, pc)
	}
}
