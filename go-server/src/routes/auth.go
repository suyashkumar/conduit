package routes

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"models"
	"net/http"
	"time"
)

const SecretKey = "lolsowow13333ksnvaa"

func returnHash(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword)
}

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

func New(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("homeauto").C("users")
	err = c.Insert(&models.User{Email: "me@suyash.io", Password: returnHash("sk1234")})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "DONE")
}

type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("homeauto").C("users")

	candidate := models.User{}
	c.Find(bson.M{"email": u.Email}).One(&candidate)
	berr := bcrypt.CompareHashAndPassword([]byte(candidate.Password), []byte(u.Password))
	if berr != nil {
		fmt.Fprintf(w, "ERR")
	} else {
		claims := MyCustomClaims{
			candidate.Email,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
				Issuer:    "homeauto",
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, jerr := token.SignedString([]byte(SecretKey))
		if jerr != nil {
			panic(jerr)
		}
		fmt.Fprintf(w, tokenString)

		fmt.Fprintf(w, "SCORE")
	}
}
