package main

import (
	"github.com/Josiah5x/Todo-List-API/config"
	"github.com/Josiah5x/Todo-List-API/controller"
	"github.com/Josiah5x/Todo-List-API/middlewares"
	"github.com/Josiah5x/Todo-List-API/pkg/mongodb"
	mongodbd "github.com/Josiah5x/Todo-List-API/platform/mongodb_platform"
	"github.com/Josiah5x/Todo-List-API/services"
	echo "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")
	todo := v1.Group("/todos", middlewares.Auths())
	todo.POST("", control.AddTodo)
	todo.PUT("/:id", control.UpdateTodo)
	todo.GET("", control.GetAllTodo)
	todo.GET("/:id", control.GetTodo)
	todo.PATCH("/mark/:id", control.MarkTodoUpdateDone)
	todo.PATCH("/deadline/:id", control.ChangeTodoDeadline)
	todo.DELETE("/delete/:id", control.DeleteTodo)
	//User Section
	// v2 := e.Group("/v2")
	user := v1.Group("/user")
	user.POST("", control.AddUser)
	user.POST("/login", control.LoginUser)
	user.PUT("/:id", control.EditUser)
	user.GET("/all", control.GetAllUser)
	user.GET("/:id", control.GetUser)
	user.PATCH("/changepassword/:id", control.ChangePassword)
	// Server Section
	e.Logger.Fatal(e.Start(":8080"))

}
