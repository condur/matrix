package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Check - HTTP OK (200) status
func Check(c *gin.Context) {
	c.Status(http.StatusOK)
}

// Pong - HTTP OK (200) status with 'pong' message
func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
