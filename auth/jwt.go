package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/Josiah5x/Todo-List-API/config"
	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	jwt.StandardClaims
}

func GenerateJWT(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	claims := &JWTClaim{
		UserName: user.UserName,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err := token.SignedString([]byte(config.Config("JWT_SCRETE_KEY")))
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, map[string]string{"token": tk}, " ")
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config("JWT_SCRETE_KEY")), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return

}
