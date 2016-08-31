package routes

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"models"
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("homeauto").C("users")

	var results []models.User
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
	resBytes, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(resBytes))
}

func Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u := models.User{}
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&u)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("sk123"), 10)
	fmt.Println(string(hashedPassword))
	fmt.Println(u.Password)
	fmt.Println(u)
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(u.Password))
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintf(w, "SCORE")
	}

}
