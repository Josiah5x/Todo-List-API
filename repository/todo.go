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
}
