package middleware

import (
	"api/logging"
	"strings"

	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {

	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "metrics") {
			return
		}
		logging.Trace(nil, "foo: bar, status: 200")
	}
}
