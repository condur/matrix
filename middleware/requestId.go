package middleware

import (
	"github.com/condur/matrix/request"
	"github.com/gin-gonic/gin"
)

// RequestId generates a request identity and save it in the header
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the request id from header
		requestId := c.Request.Header.Get(request.RequestId)

		// Generate a request id identity if needed
		if requestId == "" {
			requestId = request.NewId()
		}

		// Set the request identity in the header
		c.Writer.Header().Set(request.RequestId, requestId)

		// Process request
		c.Next()
	}
}
