package middleware

import (
	"api/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GinLogger() gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := util.GetDurationInMillseconds(start)
		if strings.Contains(c.Request.URL.Path, "metrics") || strings.Contains(c.Request.URL.Path, "health") {
			return
		}
		// wip
		log.Debug().
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Float64("duration", duration).
			Send()
	}
}
