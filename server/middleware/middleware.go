package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/conduit/server/handlers"
	"net/http"
)

func ConduitMiddleware(next handlers.ConduitHandler, c *handlers.HandlerContext) httprouter.Handle {
	middleware := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		next(w, r, ps, c)
	}

	return middleware
}

func ConduitAuthMiddleware(next handlers.AuthHandler, c *handlers.HandlerContext) httprouter.Handle {

	middleware := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.SetCorsHeaders(w)
		if candidateToken, ok := r.Header["X-Access-Token"]; ok {
			// Parse and validate token:
			token, err := jwt.ParseWithClaims(candidateToken[0], &handlers.HomeAutoClaims{}, func(token *jwt.Token) (interface{}, error) {
				return handlers.SecretKey, nil
			})

			if claims, ok := token.Claims.(*handlers.HomeAutoClaims); ok && token.Valid {
				next(w, r, ps, c, claims)
				return
			} else {
				handlers.SendErrorResponse(w, err.Error(), 401)
				fmt.Println("Error in Auth middleware")
				fmt.Println(err.Error())
				return
			}
		}
		// Either token wasn't valid or it wasn't provided
		handlers.SendErrorResponse(w, "No Token", 400)
		return
	}

	return middleware

}
