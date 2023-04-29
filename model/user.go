package model

import "time"

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
