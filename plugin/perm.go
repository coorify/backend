package plugin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/perm"
	"github.com/gin-gonic/gin"
)

func Perm(s Server) gin.HandlerFunc {
	pc := perm.NewPermCache()

	return func(c *gin.Context) {
		c.Set(field.SYS_PERMCACHE, pc)
	}
}
