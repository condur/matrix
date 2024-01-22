package middleware

import (
	"net/http"

	"github.com/condur/matrix/log"
	"github.com/gin-gonic/gin"
)

// Recovery - middleware that recovers from any panics and writes a 500 if there was one.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {

		// defer a recovery from any panic
		defer recovery(c)

		// Process request
		c.Next()
	}
}

func recovery(c *gin.Context) {
	if err := recover(); err != nil {
		// Log the error message
		log.Errorf("recovered from panic: %v", err)

		// Return a server error response
		c.Status(http.StatusInternalServerError)
	}
}
