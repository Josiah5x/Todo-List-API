package controller

import (
	"net/http"

	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/Josiah5x/Todo-List-API/services"

	"github.com/Josiah5x/Todo-List-API/pkg/structs"

	"github.com/labstack/echo"
)

type Controller struct {
	appSvc *services.Service
}

func New(svc *services.Service) *Controller {
	return &Controller{svc}
}

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

func (control *Controller) DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	err := control.appSvc.DeleteTodo(id)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, "Todo deleted Successfully", " ")
}

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
