package middlewares

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var cores = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
})

func CORESConfig() echo.MiddlewareFunc {
	return cores

}
