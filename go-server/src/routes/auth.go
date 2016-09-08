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

var SecretKey = []byte("lolsowow13333ksnvaa")

type HomeAutoClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type TokenResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type AuthHandler func(
	http.ResponseWriter,
	*http.Request,
	httprouter.Params,
	*HomeAutoClaims,
)

func AuthMiddlewareGenerator(next AuthHandler) httprouter.Handle {

	middleware := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		SetCorsHeaders(w)
		if candidateToken, ok := r.Header["X-Access-Token"]; ok {
			// Parse and validate token:
			token, err := jwt.ParseWithClaims(candidateToken[0], &HomeAutoClaims{}, func(token *jwt.Token) (interface{}, error) {
				return SecretKey, nil
			})

			if claims, ok := token.Claims.(*HomeAutoClaims); ok && token.Valid {
				next(w, r, ps, claims)
				return
			} else {
				returnError := ErrorResponse{Success: false, Error: err.Error()}
				resBytes, _ := json.Marshal(returnError)
				w.WriteHeader(401)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, string(resBytes))
				fmt.Println("there is error")
				fmt.Println(err.Error())
				return
			}
		}
		// Either token wasn't valid or it wasn't provided
		resBytes, _ := json.Marshal(ErrorResponse{
			Success: false,
			Error:   "No Token",
		})
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(resBytes))
		return
	}

	return middleware

}

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

func Test(w http.ResponseWriter, r *http.Request, ps httprouter.Params, hc *HomeAutoClaims) {
	fmt.Fprintf(w, "You're authenticated\n")
	fmt.Fprintf(w, hc.Email)
}

func Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	SetCorsHeaders(w)
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
		resBytes, _ := json.Marshal(ErrorResponse{Success: false, Error: berr.Error()})
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(resBytes))
		return
	} else {
		claims := HomeAutoClaims{
			candidate.Email,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
				Issuer:    "homeauto",
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, jerr := token.SignedString(SecretKey)
		if jerr != nil {
			panic(jerr)
		}
		w.Header().Set("Content-Type", "application/json")
		resBytes, _ := json.Marshal(TokenResponse{Success: true, Token: tokenString})
		fmt.Fprintf(w, string(resBytes))
	}
}
