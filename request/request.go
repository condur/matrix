package request

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	Id            string = "X-Request-Id"
	forwardedHost string = "X-Forwarded-Host"
)

func init() {
	uuid.EnableRandPool()
}

// NewId - generates a request id
func NewId() string {
	return uuid.NewString()
}

func GetReaderId(c *gin.Context) string {
	return c.Request.Header.Get(Id)
}

func GetWriterId(c *gin.Context) string {
	return c.Writer.Header().Get(Id)
}

func SetId(c *gin.Context, value string) {
	c.Writer.Header().Set(Id, value)
}

func Status(c *gin.Context) int {
	return c.Writer.Status()
}

func IP(c *gin.Context) string {
	return c.ClientIP()
}

func Method(c *gin.Context) string {
	return c.Request.Method
}

func URI(c *gin.Context) string {
	return c.Request.URL.RequestURI()
}

func UserAgent(c *gin.Context) string {
	return c.Request.UserAgent()
}

func ForwardedHost(c *gin.Context) string {
	return c.Request.Header.Get(forwardedHost)
}

func Error(c *gin.Context) string {
	return c.Errors.String()
}
