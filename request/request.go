package request

import (
	"github.com/google/uuid"
)

const (
	RequestId string = "X-Request-Id"
)

func init() {
	uuid.EnableRandPool()
}

// NewId - generates a request id
func NewId() string {
	return uuid.NewString()
}
