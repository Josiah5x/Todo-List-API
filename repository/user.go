package repository

import "github.com/Josiah5x/Todo-List-API/model"

type UserRepository interface {
	// Section for User
	AddUser(user *model.User) (*model.User, error)
	EditUser(user *model.User) (*model.User, error)
	GetAllUser() ([]*model.User, error)
	GetUser(id string) (*model.User, error)
	LoginUser(username string) (*model.User, error)
	ChangePassword(user *model.User) (*model.User, error)
	DeleteUser(id string) error
}
