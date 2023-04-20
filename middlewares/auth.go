package middlewares

import (
	"github.com/Josiah5x/Todo-List-API/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Auths() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.Config("JWT_SCRETE_KEY")),
	})
}
