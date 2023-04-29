package services

import (
	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/Josiah5x/Todo-List-API/repository"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodo(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo}
}

func (ss *TodoService) AddTodo(todo *model.Todo) (*model.Todo, error) {
	return ss.repo.AddTodo(todo)

}
func (ss *TodoService) EditTodo(todo *model.Todo) (*model.Todo, error) {
	return ss.repo.EditTodo(todo)

}
func (ss *TodoService) GetAllTodo() ([]*model.Todo, error) {
	return ss.repo.GetAllTodo()
}
func (ss *TodoService) GetTodo(id string) (*model.Todo, error) {
	return ss.repo.GetTodo(id)

}

func (ss *TodoService) MarkTodoDone(todo *model.Todo) (*model.Todo, error) {
	return ss.repo.MarkTodoDone(todo)

}
func (ss *TodoService) DeleteTodo(id string) error {
	return ss.repo.DeleteTodo(id)
}
func (ss *TodoService) ChangeTodoDeadline(todo *model.Todo) (*model.Todo, error) {
	return ss.repo.ChangeTodoDeadline(todo)
}
