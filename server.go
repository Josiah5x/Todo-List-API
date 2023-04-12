package main

import (
	"github.com/Josiah5x/Todo-List-API/config"
	"github.com/Josiah5x/Todo-List-API/controller"
	"github.com/Josiah5x/Todo-List-API/pkg/mongodb"
	mongodbd "github.com/Josiah5x/Todo-List-API/platform/mongodb_platform"
	"github.com/Josiah5x/Todo-List-API/services"
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

	todo.POST("", control.AddTodo)
	todo.PUT(":id", control.UpdateTodo)
	todo.GET("", control.GetAllTodo)
	todo.GET(":id", control.GetTodo)
	todo.PATCH("/mark/:id", control.MarkTodoUpdateDone)
	todo.PATCH("/deadline/:id", control.ChangeTodoDeadline)
	todo.DELETE("/delete/:id", control.DeleteTodo)
	e.Logger.Fatal(e.Start(":8080"))

}
