package mongodbd

import (
	"fmt"
	"time"

	bcryptionpassword "github.com/Josiah5x/Todo-List-API/BcryptionPassword"
	"github.com/Josiah5x/Todo-List-API/model"
	"github.com/Josiah5x/Todo-List-API/pkg/id"

	"gopkg.in/mgo.v2/bson"
)

const UserCollection = "user"

// User platform
func (mgo *Mongodb) AddUser(user *model.User) (*model.User, error) {
	user.UserId = id.GenerateNewUniqueCode()
	user.CreatedAts = time.Now()
	value, errs := bcryptionpassword.HashPassword(user.Password)
	if errs != nil {
		fmt.Println(errs)
	}
	user.Password = value
	err := mgo.db.C(UserCollection).Insert(user)
	if err != nil {
		panic(err)
	}
	return user, nil
}

// Edit a Specific User By Id
func (mgo *Mongodb) EditUser(user *model.User) (*model.User, error) {
	user.UpdatedAts = time.Now()
	err := mgo.db.C(UserCollection).Update(bson.M{"_id": user.UserId}, user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Get all register user in the Database
func (mgo *Mongodb) GetAllUser() ([]*model.User, error) {
	var user []*model.User
	err := mgo.db.C(UserCollection).Find(bson.M{}).All(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Get a Specific user by ID
func (mgo *Mongodb) GetUser(id string) (*model.User, error) {
	var user = new(model.User)
	err := mgo.db.C(UserCollection).Find(bson.M{"userid": id}).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// This where user Login by registration completed
func (mgo *Mongodb) LoginUser(username string) (*model.User, error) {
	var user = new(model.User)
	err := mgo.db.C(UserCollection).Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return user, nil
	}
	return user, nil
}

// usser can change password if needed
func (mgo *Mongodb) ChangePassword(user *model.User) (*model.User, error) {
	user.UpdatedAts = time.Now()
	err := mgo.db.C(UserCollection).Update(nil, bson.M{"$set": bson.M{"password": user.Password, "b": true}})
	if err != nil {
		return user, err
	}
	return user, nil
}

func (mgo *Mongodb) DeleteUser(id string) error {
	err := mgo.db.C(UserCollection).Remove(bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
