package repository

import "github.com/Josiah5x/Todo-List-API/model"

type Repository interface {
	AddTodo(todo *model.Todo) (*model.Todo, error)
	EditTodo(todo *model.Todo) (*model.Todo, error)
	GetAllTodo() ([]*model.Todo, error)
	GetTodo(id string) (*model.Todo, error)
	MarkTodoDone(todo *model.Todo) (*model.Todo, error)
	ChangeTodoDeadline(todo *model.Todo) (*model.Todo, error)
	DeleteTodo(id string) error
	// Section for User
	AddUser(user *model.User) (*model.User, error)
	EditUser(user *model.User) (*model.User, error)
	GetAllUser() ([]*model.User, error)
	GetUser(id string) (*model.User, error)
	LoginUser(username, password string) (*model.User, error)
	ChangePassword(user *model.User) (*model.User, error)
}
