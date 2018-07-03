package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"
	"time"
)

var client *statsd.Client
var userOptions Options

var handlerFunc = func(c *gin.Context) {
	startTime := time.Now()
	c.Next()

	if client != nil {
		printLog("Preparing metrics", infoLevel)
		// get the statsd metric key prefix provided in gin context key
		key, _ := c.Get(userOptions.getRequestKey())
		if key == nil {
			key = ""
		}

		// send status code
		status := c.Writer.Status()
		client.Increment(fmt.Sprintf("%sstatus_code.%d", key, status))

		// send response time
		duration := time.Since(startTime).Seconds() * 1000 // in milliseconds
		client.Timing(fmt.Sprintf("%sresponse_time", key), duration)

		printLog("Metrics sent", infoLevel)
	}
}

// New will setup middleware and return handler
func New(opts Options) gin.HandlerFunc {
	userOptions = opts
	addr := userOptions.getAddress()
	var clientErr error
	client, clientErr = statsd.New(statsd.Address(addr))
	if clientErr != nil {
		client = nil
		printLog(fmt.Sprintf("Failed connecting to statsd - %s", clientErr.Error()), errorLevel)
	} else {
		printLog("Sucessfully connected to statsd", infoLevel)
	}

	return handlerFunc
}
