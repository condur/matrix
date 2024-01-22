package middleware

import (
	"github.com/gin-gonic/gin"
)

const (
	cacheControlKey   string = "Cache-Control"
	cacheControlValue string = "no-cache, no-store, must-revalidate"

	pragmaKey   string = "Pragma"
	pragmaValue string = "no-cache"

	expiresKey   string = "Expires"
	expiresValue string = "0"
)

// CacheControl set the cache configuration
func CacheControl() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Set the cache for HTTP 1.1.
		c.Writer.Header().Set(cacheControlKey, cacheControlValue)

		// Set the cache for HTTP 1.0.
		c.Writer.Header().Set(pragmaKey, pragmaValue)

		// Set the cache for Proxies.
		c.Writer.Header().Set(expiresKey, expiresValue)

		// Process request
		c.Next()
	}
}
