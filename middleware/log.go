package middleware

import (
	"time"

	"github.com/condur/matrix/log"
	"github.com/condur/matrix/request"
	"github.com/gin-gonic/gin"
)

const (
	keyClientIp      string = "client_ip"
	keyDuration      string = "duration"
	keyMethod        string = "method"
	keyPath          string = "path"
	keyStatus        string = "status"
	keyUserAgent     string = "user_agent"
	keyRequestId     string = "req_id"
	keyForwardedHost string = "forwarded_host"
)

// Logger logs context HTTP requests in JSON format, with some additional custom key/values
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// End timer
		duration := time.Since(start).String()

		// Get fields to be logged
		fields := []log.Field{
			log.String(keyDuration, duration),
			log.String(keyClientIp, request.IP(c)),
			log.String(keyMethod, request.Method(c)),
			log.String(keyPath, request.URI(c)),
			log.Int(keyStatus, request.Status(c)),
			log.String(keyUserAgent, request.UserAgent(c)),
			log.String(keyRequestId, request.GetWriterId(c)),
		}

		// Log the Forwarded Host, if present
		if forwardedHost := request.ForwardedHost(c); forwardedHost != "" {
			fields = append(fields, log.String(keyForwardedHost, forwardedHost))
		}

		// Log the request status
		if request.Status(c) >= 400 {
			log.Error(request.Error(c), fields...)
		} else {
			log.Info("ok", fields...)
		}
	}
}
