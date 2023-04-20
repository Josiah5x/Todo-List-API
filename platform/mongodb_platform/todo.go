package mongodbd

import (
	"time"

	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/Josiah5x/Todo-List-API/pkg/id"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongodb struct {
	db *mgo.Database
}

func New(db *mgo.Database) *Mongodb {
	return &Mongodb{db}

}

func (mgo *Mongodb) AddTodo(todo *model.Todo) (*model.Todo, error) {
	todo.Id = id.GenerateNewUniqueCode()
	todo.Status = "Undone"
	todo.CreatedAt = time.Now()
	err := mgo.db.C("todos").Insert(todo)
	if err != nil {
		panic(err)
	}
	return todo, nil
}
func (mgo *Mongodb) EditTodo(todo *model.Todo) (*model.Todo, error) {
	todo.UpdatedAt = time.Now()
	err := mgo.db.C("todos").Update(bson.M{"_id": todo.Id}, todo)
	if err != nil {
		return todo, err
	}
	return todo, nil

}
func (mgo *Mongodb) GetTodo(id string) (*model.Todo, error) {
	var todo = new(model.Todo)
	err := mgo.db.C("todos").Find(bson.M{"_id": id}).One(&todo)
	if err != nil {
		return todo, err
	}
	return todo, nil

}
func (mgo *Mongodb) GetAllTodo() ([]*model.Todo, error) {
	var todo []*model.Todo
	err := mgo.db.C("todos").Find(bson.M{}).All(&todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (mgo *Mongodb) MarkTodoDone(todo *model.Todo) (*model.Todo, error) {
	err := mgo.db.C("todos").Update(nil, bson.M{"$set": bson.M{"status": todo.Status, "b": true}})
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (mgo *Mongodb) ChangeTodoDeadline(todo *model.Todo) (*model.Todo, error) {
	err := mgo.db.C("todos").Update(nil, bson.M{"$set": bson.M{"deadline": todo.Deadline, "b": true}})
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (mgo *Mongodb) DeleteTodo(id string) error {
	err := mgo.db.C("todos").Remove(bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

// User platform
func (mgo *Mongodb) AddUser(user *model.User) (*model.User, error) {
	user.UserId = id.GenerateNewUniqueCode()
	user.CreatedAts = time.Now()
	err := mgo.db.C("user").Insert(user)
	if err != nil {
		panic(err)
	}
	return user, nil
}

// Edit a Specific User By Id
func (mgo *Mongodb) EditUser(user *model.User) (*model.User, error) {
	user.UpdatedAts = time.Now()
	err := mgo.db.C("user").Update(bson.M{"_id": user.UserId}, user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Get all register user in the Database
func (mgo *Mongodb) GetAllUser() ([]*model.User, error) {
	var user []*model.User
	err := mgo.db.C("user").Find(bson.M{}).All(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Get a Specific user by ID
func (mgo *Mongodb) GetUser(id string) (*model.User, error) {
	var user = new(model.User)
	err := mgo.db.C("user").Find(bson.M{"_id": id}).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// This where user Login by registration completed
func (mgo *Mongodb) LoginUser(username, password string) (*model.User, error) {
	var user = new(model.User)
	err := mgo.db.C("user").Find(bson.M{"username": username, "password": password}).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// usser can change password if needed
func (mgo *Mongodb) ChangePassword(user *model.User) (*model.User, error) {
	user.UpdatedAts = time.Now()
	err := mgo.db.C("user").Update(nil, bson.M{"$set": bson.M{"password": user.Password, "b": true}})
	if err != nil {
		return user, err
	}
	return user, nil
}
