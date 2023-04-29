package model

import "time"

type Todo struct {
	Id           string    `json:"_id" bson:"_id"`
	UserId       string    `json:"userid" bson:"userid"`
	Name         string    `json:"name" bson:"name"`
	Descriptions string    `json:"description" bson:"description"`
	Status       string    `json:"status" bson:"status"`
	Deadline     string    `json:"deadline" bson:"deadline"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}
