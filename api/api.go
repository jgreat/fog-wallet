//go:generate oapi-codegen --config=types.cfg.yaml ../openapi/fog-wallet.yaml
//go:generate oapi-codegen --config=server.cfg.yaml ../openapi/fog-wallet.yaml

package api

import (
	"log"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Api -
type Api struct {
	DB *gorm.DB
}

// NewApi - initialize Api
func New() *Api {
	return &Api{}
}

/**
ValidationMiddleware - setup swagger document validation.
Validation to check all requests against the OpenAPI schema.

If paths don't exist in validation returns:
	"no matching operation was found"
**/
func ValidationMiddleware() echo.MiddlewareFunc {
	// swagger spec is embedded in the generated code.
	swagger, err := GetSwagger()
	if err != nil {
		log.Fatalf("Error loading swagger spec\n: %s", err)
	}

	/**
	Clear out the servers array in the swagger spec, that skips validating
	that server names match. We don't know how this thing will be run.
	**/
	swagger.Servers = nil

	return middleware.OapiRequestValidator(swagger)
}

// Send an echo server error
func sendApiError(ctx echo.Context, code int, message string) error {
	apiErr := Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, apiErr)
	return err
}
