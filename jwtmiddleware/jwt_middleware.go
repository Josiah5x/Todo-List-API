package jwtmiddleware

import (
	"github.com/Josiah5x/Todo-List-API/config"
	"github.com/labstack/echo/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(config.Config("JWT_SCRETE_KEY")),
})
