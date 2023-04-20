package controller

import (
	"fmt"
	"net/http"

	"github.com/Josiah5x/Todo-List-API/auth"
	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/Josiah5x/Todo-List-API/pkg/structs"
	"github.com/Josiah5x/Todo-List-API/services"
	"github.com/labstack/echo"
)

// type jwtCustomClaims struct {
// 	FirstName string `json:"fname"`
// 	LastName  string `json:"lname"`
// 	// Admin     bool   `json:"admin"`
// 	jwt.StandardClaims
// }

type Controller struct {
	appSvc *services.Service
}

func New(svc *services.Service) *Controller {
	return &Controller{svc}
}

// Add or Create Todo in the Database
func (control *Controller) AddTodo(c echo.Context) error {
	newTodo := new(model.Todo)
	if err := c.Bind(newTodo); err != nil {
		return err
	}
	todo, err := control.appSvc.AddTodo(newTodo)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, todo, " ")
}

// Update a Todo in the Database
func (control *Controller) UpdateTodo(c echo.Context) error {
	newtodo := new(model.Todo)

	id := string(c.Param("id"))
	if id == "" {
		return nil
	}
	if err := c.Bind(newtodo); err != nil {
		return err
	}
	findEm, err := control.appSvc.GetTodo(id)
	if err != nil {
		return err
	}
	structs.Merge(findEm, newtodo)
	todo, err := control.appSvc.EditTodo(findEm)
	// fmt.Println("todo", todo)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, todo, " ")
}

// Return all Todo in the Database
func (control *Controller) GetAllTodo(c echo.Context) error {
	todo, err := control.appSvc.GetAllTodo()
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, todo, " ")

}
func (control *Controller) GetTodo(c echo.Context) error {
	id := c.Param("id")
	todo, err := control.appSvc.GetTodo(id)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, todo, " ")

}

// Delete Todo By ID
func (control *Controller) DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	err := control.appSvc.DeleteTodo(id)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, "Todo deleted Successfully", " ")
}

// Mark Todo eg(Done/Undone) By ID
func (control *Controller) MarkTodoUpdateDone(c echo.Context) error {
	newtodo := new(model.Todo)

	id := string(c.Param("id"))
	if id == "" {
		return nil
	}
	if err := c.Bind(newtodo); err != nil {
		return err
	}
	findEm, err := control.appSvc.GetTodo(id)
	// fmt.Println(findEm)
	if err != nil {
		return err
	}
	structs.Merge(findEm, newtodo)
	todo, err := control.appSvc.MarkTodoDone(findEm)
	// fmt.Println(todo)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, todo, " ")
}

// Change a Specific todos Deadline by ID
func (control *Controller) ChangeTodoDeadline(c echo.Context) error {
	newtodo := new(model.Todo)

	id := string(c.Param("id"))
	if id == "" {
		return nil
	}
	if err := c.Bind(newtodo); err != nil {
		return err
	}
	findEm, err := control.appSvc.GetTodo(id)
	// fmt.Println(findEm)
	if err != nil {
		return err
	}
	structs.Merge(findEm, newtodo)
	todo, err := control.appSvc.ChangeTodoDeadline(findEm)
	// fmt.Println(todo)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, todo, " ")
}

// Add or create a user
func (control *Controller) AddUser(c echo.Context) error {
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
func (control *Controller) EditUser(c echo.Context) error {
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
func (control *Controller) GetAllUser(c echo.Context) error {
	user, err := control.appSvc.GetAllUser()
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, user, " ")

}

// Return a  Specific user by Id
func (control *Controller) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := control.appSvc.GetUser(id)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, user, " ")
}

//  Login for Auth using JSON Web Token
func (control *Controller) LoginUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	findUser, err := control.appSvc.LoginUser(username, password)
	if err != nil {
		return err
	}
	// fmt.Println(findUser.UserName, findUser.Password)
	if findUser.UserName == username && findUser.Password == password {
		auth.GenerateJWT(c)
	}
	return echo.ErrUnauthorized
}

// Change user password, by id
func (control *Controller) ChangePassword(c echo.Context) error {
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
