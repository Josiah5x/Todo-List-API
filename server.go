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
	svc := services.NewTodo(mgos)
	svcs := services.NewUser(mgos)
	// Controller
	Todocontrol := controller.NewTodo(svc)
	Usercontrol := controller.NewUser(svcs)
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.CORESConfig())

	v1 := e.Group("/v1")
	todo := v1.Group("/todos", middlewares.Auths())
	todo.POST("", Todocontrol.AddTodo)
	todo.PUT("/:id", Todocontrol.UpdateTodo)
	todo.GET("", Todocontrol.GetAllTodo)
	todo.GET("/id", Todocontrol.GetTodo)
	todo.PATCH("/mark/:id", Todocontrol.MarkTodoUpdateDone)
	todo.PATCH("/deadline/:id", Todocontrol.ChangeTodoDeadline)
	todo.DELETE("/delete/:id", Todocontrol.DeleteTodo)
	//User Section
	// v2 := e.Group("/v2")
	user := v1.Group("/user")
	user.POST("", Usercontrol.AddUser)
	user.POST("/login", Usercontrol.LoginUser)
	user.PUT("/:id", Usercontrol.EditUser)
	user.GET("/all", Usercontrol.GetAllUser)
	user.GET("/:id", Usercontrol.GetUser)
	user.PATCH("/changepassword/:id", Usercontrol.ChangePassword)
	todo.DELETE("/delete/:id", Usercontrol.DeleteUser)
	// Server Section
	e.Logger.Fatal(e.Start(":8080"))

}
