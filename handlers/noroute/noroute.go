package noroute

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NoRoute - HTTP NotFound (404) status
func NoRoute(c *gin.Context) {
	c.Status(http.StatusNotFound)
}
