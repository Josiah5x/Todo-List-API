package services

import (
	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/Josiah5x/Todo-List-API/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo}
}

func (ss *Service) AddTodo(todo *model.Todo) (*model.Todo, error) {
	return ss.repo.AddTodo(todo)

}
func (ss *Service) EditTodo(todo *model.Todo) (*model.Todo, error) {
	return ss.repo.EditTodo(todo)

}
func (ss *Service) GetAllTodo() ([]*model.Todo, error) {
	return ss.repo.GetAllTodo()
}
func (ss *Service) GetTodo(id string) (*model.Todo, error) {
	return ss.repo.GetTodo(id)

}

func (ss *Service) MarkTodoDone(todo *model.Todo) (*model.Todo, error) {
	return ss.repo.MarkTodoDone(todo)

}
func (ss *Service) DeleteTodo(id string) error {
	return ss.repo.DeleteTodo(id)
}
func (ss *Service) ChangeTodoDeadline(todo *model.Todo) (*model.Todo, error) {
	return ss.repo.ChangeTodoDeadline(todo)
}

// User Services

func (ss *Service) AddUser(user *model.User) (*model.User, error) {
	return ss.repo.AddUser(user)

}
func (ss *Service) EditUser(user *model.User) (*model.User, error) {
	return ss.repo.EditUser(user)
}
func (ss *Service) GetAllUser() ([]*model.User, error) {
	return ss.repo.GetAllUser()
}
func (ss *Service) GetUser(id string) (*model.User, error) {
	return ss.repo.GetUser(id)
}
func (ss *Service) LoginUser(username, password string) (*model.User, error) {
	return ss.repo.LoginUser(username, password)
}

func (ss *Service) ChangePassword(user *model.User) (*model.User, error) {
	return ss.repo.ChangePassword(user)
}
