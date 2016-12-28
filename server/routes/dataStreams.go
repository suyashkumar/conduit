package routes

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/conduit/server/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func GetStreamedMessages(w http.ResponseWriter, r *http.Request, ps httprouter.Params, hc *HomeAutoClaims) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("homeauto").C("streammessages")

	var results []models.StreamMessage
	prefixedName := PrefixedName(ps.ByName("deviceName"), hc.Prefix)
	topicName := prefixedName + "/stream/" + ps.ByName("streamName")
	err = c.Find(bson.M{"topic": topicName}).All(&results)
	if err != nil {
		panic(err)
	}
	resBytes, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(resBytes))
}
