package plugin

import (
	"github.com/coorify/backend/field"
	"github.com/gin-gonic/gin"
)

func Option(s Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(field.SYS_OPTION, s.Option())
	}
}
