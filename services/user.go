package services

import (
	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/Josiah5x/Todo-List-API/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUser(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

// User Services

func (ss *UserService) AddUser(user *model.User) (*model.User, error) {
	return ss.repo.AddUser(user)

}
func (ss *UserService) EditUser(user *model.User) (*model.User, error) {
	return ss.repo.EditUser(user)
}
func (ss *UserService) GetAllUser() ([]*model.User, error) {
	return ss.repo.GetAllUser()
}
func (ss *UserService) GetUser(id string) (*model.User, error) {
	return ss.repo.GetUser(id)
}

func (ss *UserService) LoginUser(username string) (*model.User, error) {
	return ss.repo.LoginUser(username)
}

func (ss *UserService) ChangePassword(user *model.User) (*model.User, error) {
	return ss.repo.ChangePassword(user)
}

func (ss *UserService) DeleteUser(id string) error {
	return ss.repo.DeleteUser(id)
}
