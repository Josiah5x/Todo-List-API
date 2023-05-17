package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/Josiah5x/Todo-List-API/pkg/structs"
	"github.com/Josiah5x/Todo-List-API/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type TodoController struct {
	appSvc *services.TodoService
}

func NewTodo(svc *services.TodoService) *TodoController {
	return &TodoController{svc}
}

// Add or Create Todo in the Database
func (control *TodoController) AddTodo(c echo.Context) error {
	newtodo := new(model.Todo)
	if err := c.Bind(newtodo); err != nil {
		return err
	}
	// ....................
	user := c.Get("user").(*jwt.Token)
	// fmt.Println("see what ar looking 4", user)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["UserId"].(string)
	// fmt.Println("THIS IS THE ID:", userid)
	newtodo.UserId = userid
	// time Duration for Deadline
	duration, err := time.ParseDuration(newtodo.Deadline)
	if err != nil {
		panic(err)
	}
	deadline := time.Now().Local().Add(duration)
	newtodo.Deadline = deadline.Format("2 Jan 06 03:04PM")
	// .....................
	_, err = control.appSvc.AddTodo(newtodo)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, "Successfully Added", " ")
}

// Update a Todo in the Database
func (control *TodoController) UpdateTodo(c echo.Context) error {
	newtodo := new(model.Todo)

	id := string(c.Param("id"))
	if id == "" {
		return nil
	}
	if err := c.Bind(newtodo); err != nil {
		return err
	}
	user := c.Get("user").(*jwt.Token)
	// fmt.Println("see what ar looking 4", user)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["UserId"].(string)
	// fmt.Println("THIS IS THE ID:", userid)
	newtodo.UserId = userid

	duration, err := time.ParseDuration(newtodo.Deadline)
	if err != nil {
		panic(err)
	}
	deadline := time.Now().Local().Add(duration)
	newtodo.Deadline = deadline.Format("2 Jan 06 03:04PM")
	findEm, err := control.appSvc.GetTodo(id)
	if err != nil {
		return err
	}
	structs.Merge(findEm, newtodo)
	_, err = control.appSvc.EditTodo(findEm)
	// fmt.Println("todo", todo)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, "Update Successfully", " ")
}

// Return all Todo in the Database
func (control *TodoController) GetAllTodo(c echo.Context) error {
	var newTodoList []*model.Todo
	// id := c.Param("id")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["UserId"].(string)
	//.......................................
	todo, err := control.appSvc.GetAllTodo()
	if err != nil {
		return err
	}
	for _, v := range todo {
		if v.UserId == userid {
			newTodoList = append(newTodoList, v)
		}
	}
	return c.JSONPretty(http.StatusOK, newTodoList, " ")

}
func (control *TodoController) GetTodo(c echo.Context) error {
	id := c.Param("id")
	todo, err := control.appSvc.GetTodo(id)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, todo, " ")

}

// Mark Todo eg(Done/Undone) By ID
func (control *TodoController) MarkTodoUpdateDone(c echo.Context) error {
	newtodo := new(model.Todo)
	id := string(c.Param("id"))
	if id == "" {
		return nil
	}
	if err := c.Bind(newtodo); err != nil {
		return err
	}
	findEm, err := control.appSvc.GetTodo(id)
	fmt.Println(findEm)
	if err != nil {
		return err
	}
	structs.Merge(findEm, newtodo)
	_, err = control.appSvc.MarkTodoDone(findEm)
	// fmt.Println(todo.Status)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, "Updated Successfully", " ")
}

// Change a Specific todos Deadline by ID
func (control *TodoController) ChangeTodoDeadline(c echo.Context) error {
	newtodo := new(model.Todo)

	id := string(c.Param("id"))
	if id == "" {
		return nil
	}
	if err := c.Bind(newtodo); err != nil {
		return err
	}
	duration, err := time.ParseDuration(newtodo.Deadline)
	if err != nil {
		panic(err)
	}
	deadline := time.Now().Local().Add(duration)
	// x, _ := time.ParseDuration(newtodo.Deadline)
	// deadline := time.Now().Add(time.Hour * x)
	newtodo.Deadline = deadline.Format("2 Jan 06 03:04PM")
	findEm, err := control.appSvc.GetTodo(id)
	// fmt.Println(findEm)
	if err != nil {
		return err
	}
	structs.Merge(findEm, newtodo)
	_, err = control.appSvc.ChangeTodoDeadline(findEm)
	// fmt.Println(todo)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, "Updated Successfully", " ")
}

// Delete Todo By ID
func (control *TodoController) DeleteTodo(c echo.Context) error {
	id := string(c.Param("id"))
	err := control.appSvc.DeleteTodo(id)
	if err != nil {
		return err
	}
	//
	return c.JSONPretty(http.StatusOK, "Deleted Successfully", " ")
}
