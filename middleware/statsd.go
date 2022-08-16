package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"

	"fmt"
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

		path := strings.ReplaceAll(c.FullPath(), "/", "_")
		path = strings.ReplaceAll(path, "*", "_")
		path = strings.ReplaceAll(path, ":", "_")

		// send status code
		status := c.Writer.Status()
		client.Increment(fmt.Sprintf("%s.%s.status_code.%d", key, path, status))

		// send response time
		duration := time.Since(startTime).Seconds() * 1000 // in milliseconds
		client.Timing(fmt.Sprintf("%s.%s.response_time", key, path), duration)

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
		printLog("Successfully connected to statsd", infoLevel)
	}

	return handlerFunc
}
