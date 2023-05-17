package auth

import (
	"errors"
	"fmt"
	"time"

	bcryptionpassword "github.com/Josiah5x/Todo-List-API/BcryptionPassword"
	"github.com/Josiah5x/Todo-List-API/config"
	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/dgrijalva/jwt-go"
)

// var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	UserName string
	UserId   string
	jwt.StandardClaims
}

func GenerateJWT(usernameForm, passwordForm string, findUser *model.User) (tk string, err error) {

	check := bcryptionpassword.CheckPassword(findUser.Password, passwordForm)
	fmt.Println(check)
	truth := true
	if usernameForm == findUser.UserName && check == truth {
		claims := &JWTClaim{
			UserName: findUser.UserName,
			UserId:   findUser.UserId,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tk, err := token.SignedString([]byte(config.Config("JWT_SCRETE_KEY")))
		if err != nil {
			return tk, err
		}
		return tk, nil
	}
	return tk, nil
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
