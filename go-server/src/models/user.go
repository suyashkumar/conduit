package models

import "gopkg.in/mgo.v2/bson"

type (
	User struct {
		Id       bson.ObjectId `json:"id" bson:"_id"`
		Email    string        `json:"email" bson:"email"`
		Password string        `json:"password" bson:"password`
	}
)
