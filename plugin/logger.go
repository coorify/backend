package plugin

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/coorify/backend/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(opt interface{}) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		path := c.Request.URL.Path
		dlen := c.Writer.Size()
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		latencyMs := float64(latency) / float64(time.Millisecond)

		fields := logrus.Fields{
			"hostname":   hostname,
			"status":     statusCode,
			"method":     method,
			"clientIP":   clientIP,
			"path":       path,
			"latency-ms": latencyMs,
			"dataLength": dlen,
		}
		entry := logger.WithFields(fields)

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%4s %3d - %s (%f)ms", method, statusCode, path, latencyMs)

			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
