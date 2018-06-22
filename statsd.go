package stats

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
