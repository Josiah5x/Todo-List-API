package main

import (
	"github.com/Josiah5x/Todo-List-API/config"
	"github.com/Josiah5x/Todo-List-API/controller"
	mongodbd "github.com/Josiah5x/Todo-List-API/platform/mongodb_platform"
	"github.com/Josiah5x/Todo-List-API/services"

	"github.com/Josiah5x/Todo-List-API/pkg/mongodb"

	echo "github.com/labstack/echo"
)

func main() {

	//connect mongodb
	db := mongodb.New(config.Config("MONGO_DB"))
	//connect Postgresqldb
	// Platform
	mgos := mongodbd.New(db)
	// Service
	svc := services.New(mgos)
	// Controller
	control := controller.New(svc)

	e := echo.New()
	// add middleware and routes
	// ...

	v1 := e.Group("/v1")

	todo := v1.Group("/todo")

	todo.POST("/addtodo", control.AddTodo)
	todo.POST("/edittodo/:id", control.UpdateTodo)
	todo.GET("/alltodo", control.GetAllTodo)
	todo.GET("/gettodo/:id", control.GetTodo)
	todo.POST("/editmark/:id", control.MarkTodoUpdateDone)
	todo.POST("/editdeadline/:id", control.ChangeTodoDeadline)
	todo.GET("/deletetodo/:id", control.DeleteTodo)
	e.Logger.Fatal(e.Start(":8080"))

}
