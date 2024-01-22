package validate

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
)

// Define and initialize validator
var validate = validator.New()

// Struct validates a structs exposed fields, and automatically validates nested structs
func Struct(ctx context.Context, s any) error {
	return validate.StructCtx(ctx, s)
}
