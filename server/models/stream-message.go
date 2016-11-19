package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	StreamMessage struct {
		Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
		Data      string        `json:"data" bson:"data"`
		Topic     string        `json:"topic" bson:"topic"`
	}
)
