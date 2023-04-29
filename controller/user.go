package controller

import (
	"fmt"
	"net/http"

	"github.com/Josiah5x/Todo-List-API/auth"
	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/Josiah5x/Todo-List-API/pkg/structs"
	"github.com/Josiah5x/Todo-List-API/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JWTClaim struct {
	UserName string
	UserId   string
	jwt.StandardClaims
}

type UserController struct {
	appSvc *services.UserService
}

func NewUser(svc *services.UserService) *UserController {
	return &UserController{svc}
}

// Add or create a user
func (control *UserController) AddUser(c echo.Context) error {
	newUser := new(model.User)
	if err := c.Bind(newUser); err != nil {
		return err
	}
	user, err := control.appSvc.AddUser(newUser)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, user, " ")
}

// Edit a User
func (control *UserController) EditUser(c echo.Context) error {
	newUser := new(model.User)
	id := string(c.Param("id"))
	if id == "" {
		return nil
	}
	if err := c.Bind(newUser); err != nil {
		return err
	}
	findEm, err := control.appSvc.GetUser(id)
	if err != nil {
		return err
	}
	structs.Merge(findEm, newUser)
	user, err := control.appSvc.EditUser(findEm)
	// fmt.Println("user", user)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, user, " ")
}

// Return All Users
func (control *UserController) GetAllUser(c echo.Context) error {
	user, err := control.appSvc.GetAllUser()
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, user, " ")

}

// Return a  Specific user by Id
func (control *UserController) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := control.appSvc.GetUser(id)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, user, " ")
}

//  Login for Auth using JSON Web Token
func (control *UserController) LoginUser(c echo.Context) error {
	usernameForm := c.FormValue("username")
	passwordForm := c.FormValue("password")

	findUser, err := control.appSvc.LoginUser(usernameForm)
	if err != nil {
		return err
	}
	tk, err := auth.GenerateJWT(usernameForm, passwordForm, findUser)
	if err != nil {
		return err
	}
	if tk == "" {

		return c.JSONPretty(http.StatusBadRequest, "Invalid Username or Password", " ")
	}
	return c.JSONPretty(http.StatusOK, map[string]string{"token": tk}, " ")
}

// Change user password, by id
func (control *UserController) ChangePassword(c echo.Context) error {
	newUser := new(model.User)
	id := string(c.Param("id"))
	// println(id)
	if id == "" {
		return nil
	}
	if err := c.Bind(newUser); err != nil {
		return err
	}
	findEm, err := control.appSvc.GetUser(id)
	// fmt.Println(findEm)
	if err != nil {
		return err
	}
	structs.Merge(findEm, newUser)
	user, err := control.appSvc.ChangePassword(findEm)
	fmt.Println(user)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, user, " ")
}

// Delete Todo By ID
func (control *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := control.appSvc.DeleteUser(id)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, "User deleted Successfully", " ")
}
