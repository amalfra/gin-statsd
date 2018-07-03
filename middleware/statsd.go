package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"
)

// New will setup middleware and return handler
func New(opts Options) gin.HandlerFunc {
	addr := opts.getAddress()
	client, clientErr := statsd.New(statsd.Address(addr))
	if clientErr != nil {
		client = nil
		printLog(fmt.Sprintf("Failed connecting to statsd - %s", clientErr.Error()), errorLevel)
	} else {
		printLog("Sucessfully connected to statsd", infoLevel)
	}

	return func(c *gin.Context) {
		c.Next()
		if client != nil {
			printLog("Preparing metrics", infoLevel)
		}
	}
}
