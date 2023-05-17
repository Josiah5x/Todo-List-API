package mongodbd

import (
	"time"

	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/Josiah5x/Todo-List-API/pkg/id"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const TodoCollection = "todos"

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
	err := mgo.db.C(TodoCollection).Insert(todo)
	if err != nil {
		panic(err)
	}
	return todo, nil
}
func (mgo *Mongodb) EditTodo(todo *model.Todo) (*model.Todo, error) {
	todo.UpdatedAt = time.Now()
	err := mgo.db.C(TodoCollection).Update(bson.M{"_id": todo.Id}, todo)
	if err != nil {
		return todo, err
	}
	return todo, nil

}
func (mgo *Mongodb) GetTodo(id string) (*model.Todo, error) {
	var todo = new(model.Todo)
	err := mgo.db.C(TodoCollection).Find(bson.M{"_id": id}).One(&todo)
	if err != nil {
		return todo, err
	}
	return todo, nil

}
func (mgo *Mongodb) GetAllTodo() ([]*model.Todo, error) {
	var todo []*model.Todo
	err := mgo.db.C(TodoCollection).Find(bson.M{}).All(&todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (mgo *Mongodb) MarkTodoDone(todo *model.Todo) (*model.Todo, error) {
	err := mgo.db.C(TodoCollection).Update(bson.M{"_id": todo.Id}, bson.M{"$set": bson.M{"status": todo.Status, "b": true}})
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (mgo *Mongodb) ChangeTodoDeadline(todo *model.Todo) (*model.Todo, error) {
	err := mgo.db.C(TodoCollection).Update(bson.M{"_id": todo.Id}, bson.M{"$set": bson.M{"deadline": todo.Deadline, "b": true}})
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (mgo *Mongodb) DeleteTodo(id string) error {
	err := mgo.db.C(TodoCollection).Remove(bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
