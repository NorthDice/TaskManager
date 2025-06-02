package model

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	Id       bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string        `bson:"username" json:"username"`
	Password string        `bson:"password" json:"password"`
}
