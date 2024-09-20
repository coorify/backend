package router

import (
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	reply.OkWithMessage("v1.3.0", c)
}
