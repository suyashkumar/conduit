package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/conduit/server/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func GetStreamedMessages(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context *HandlerContext, hc *HomeAutoClaims) {
	session := context.DbSession.New()
	defer session.Close()

	c := session.DB("homeauto").C("streammessages")

	prefixedName := PrefixedName(ps.ByName("deviceName"), hc.Prefix)
	topicName := prefixedName + "/stream/" + ps.ByName("streamName")

	var results []models.StreamMessage
	err = c.Find(bson.M{"topic": topicName}).All(&results)
	if err != nil {
		panic(err)
	}
	resBytes, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(resBytes))
}
