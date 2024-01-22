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
	keyForwardedHost string = "forward_host"
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
			log.String(keyClientIp, c.ClientIP()),
			log.String(keyMethod, c.Request.Method),
			log.String(keyPath, c.Request.URL.RequestURI()),
			log.Int(keyStatus, c.Writer.Status()),
			log.String(keyUserAgent, c.Request.UserAgent()),
			log.String(keyRequestId, c.Writer.Header().Get(request.RequestId)),
		}

		// Log the Forward Host, if present
		if forwardHost := c.Request.Header.Get("X-Forwarded-Host"); forwardHost != "" {
			fields = append(fields, log.String(keyForwardedHost, forwardHost))
		}

		// Log the information
		if c.Writer.Status() >= 400 {
			log.Error(c.Errors.String(), fields...)
		} else {
			log.Info("ok", fields...)
		}
	}
}
