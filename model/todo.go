package model

import "time"

type Todo struct {
	Id           string    `json:"_id" bson:"_id"`
	UserId       string    `json:"userid" bson:"userid"`
	Name         string    `json:"name" bson:"name"`
	Descriptions string    `json:"description" bson:"description"`
	Status       string    `json:"status" bson:"status"`
	Deadline     string    `json:"deadline" bson:"deadline"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

type User struct {
	UserId     string    `json:"_id" bson:"_id"`
	FisrName   string    `json:"firstname" bson:"firstname"`
	LastName   string    `json:"lastname" bson:"lastname"`
	Email      string    `json:"email" bson:"email"`
	UserName   string    `json:"username" bson:"username"`
	Password   string    `json:"password" bson:"password"`
	CreatedAts time.Time `bson:"created_at" json:"create_at"`
	UpdatedAts time.Time `bson:"updated_at" json:"update_at"`
}
