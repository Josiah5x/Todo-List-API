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

	v1 := e.Group("/v1")

	todo := v1.Group("/todo")

	todo.POST("/todo", control.AddTodo)
	todo.PATCH("/todo/:id", control.UpdateTodo)
	todo.GET("/todo", control.GetAllTodo)
	todo.GET("/todo/:id", control.GetTodo)
	todo.PATCH("/mark/:id", control.MarkTodoUpdateDone)
	todo.PATCH("/deadline/:id", control.ChangeTodoDeadline)
	todo.DELETE("/delete/:id", control.DeleteTodo)
	e.Logger.Fatal(e.Start(":8080"))

}
