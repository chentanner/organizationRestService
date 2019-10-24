package middleware 

import (
	"fmt"
	"os"
	api "organizationRestService/api/gen"
	
	"github.com/labstack/echo/v4"
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
)

// NewOapiValidationMiddleware creates a new oapiMiddleware with default headers.
func NewOapiValidationMiddleware() echo.MiddlewareFunc {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	return oapimiddleware.OapiRequestValidator(swagger)
}