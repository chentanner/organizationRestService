package middleware 

import (
	"os"
	"strings"
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

// NewCORSMiddleware creates a new CORSMiddleware with default headers.
func NewCORSMiddleware() echo.MiddlewareFunc {


	cors := os.Getenv("CORS") //Get port from .env file, we did not specify any port so this should return an empty string when tested locallyQ
	if cors == "" {
		cors = "*" //localhost
	}

	corsHeaders := strings.Split(cors, ",")
		if corsHeaders == nil || len(corsHeaders) == 0 {
			corsHeaders = []string{"*"}
		}

	fmt.Printf("Allowed cors: %v\n", corsHeaders)

	return echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: corsHeaders,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-Requested-With", "X-Registry-Auth"},
		AllowMethods: []string{http.MethodHead, http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete,http.MethodOptions},
	})
}