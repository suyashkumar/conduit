package mqtt

import (
	"gopkg.in/mgo.v2"
	"models"
	"time"
)

func PersistMessage(message string, topic string) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("homeauto").C("streammessages")
	err = c.Insert(&models.StreamMessage{
		Data:      message,
		Timestamp: time.Now(),
		Topic:     topic,
	})
	if err != nil {
		panic(err)
	}
}
