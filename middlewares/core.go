package middlewares

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CORESConfig() echo.MiddlewareFunc {
	var cores = middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	})
	return cores

}
