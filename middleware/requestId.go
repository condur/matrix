package middleware

import (
	"github.com/condur/matrix/request"
	"github.com/gin-gonic/gin"
)

// RequestId generates a request identity and save it in the header
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the request id from header
		id := request.GetReaderId(c)

		// Generate a request id identity if needed
		if id == "" {
			id = request.NewId()
		}

		// Set the request identity in the header
		request.SetId(c, id)

		// Process request
		c.Next()
	}
}
