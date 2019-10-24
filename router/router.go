package router

import (
	api "organizationRestService/api/gen"
	"organizationRestService/api/handlers"
	"organizationRestService/middleware"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func BuildRouter(organizationHandlers *handlers.OrganizationHandlers) *echo.Echo {

	echoRouter := echo.New()
	// Log all requests
	echoRouter.Use(echomiddleware.Logger())

	//Add CORS middleware
	echoRouter.Use(middleware.NewCORSMiddleware())

	// Validate requests against OpenAPI schema.
	echoRouter.Use(middleware.NewOapiValidationMiddleware())

	// We now register our OpenApi generated handlers
	api.RegisterHandlers(echoRouter, organizationHandlers)

	return echoRouter
}
