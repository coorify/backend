package reply

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Reply struct {
	Code    int         `json:"code"`
	Payload interface{} `json:"payload,omitempty"`
	Message string      `json:"message"`
}

func NewReply(code int, payload interface{}, message string, c *gin.Context) {
	c.JSON(code, Reply{
		Code:    code,
		Payload: payload,
		Message: message,
	})
}

func Ok(c *gin.Context) {
	NewReply(http.StatusOK, nil, "Ok", c)
}

func OkWithMessage(message string, c *gin.Context) {
	NewReply(http.StatusOK, nil, message, c)
}

func OkWithPayload(payload interface{}, c *gin.Context) {
	NewReply(http.StatusOK, payload, "Ok", c)
}

func OkWithDetailed(payload interface{}, message string, c *gin.Context) {
	NewReply(http.StatusOK, payload, message, c)
}

func Fail(c *gin.Context) {
	NewReply(http.StatusBadRequest, nil, "Fail", c)
}

func FailWithMessage(message string, c *gin.Context) {
	NewReply(http.StatusBadRequest, nil, message, c)
}

func FailWithDetailed(payload interface{}, message string, c *gin.Context) {
	NewReply(http.StatusBadRequest, payload, message, c)
}
