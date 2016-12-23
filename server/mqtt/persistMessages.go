package mqtt

import (
	"github.com/suyashkumar/conduit/server/models"
	"gopkg.in/mgo.v2"
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
