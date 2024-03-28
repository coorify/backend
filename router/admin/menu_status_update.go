package admin

import (
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/reply"
	"github.com/coorify/backend/types"
	"github.com/gin-gonic/gin"
)

func MenuStatusUpdate(c *gin.Context) {
	if err := types.StatusUpdate(c, &model.Menu{}); err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}
	reply.Ok(c)
}
