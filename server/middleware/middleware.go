package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/conduit/server/routes"
	"net/http"
)

func ConduitMiddleware(next routes.ConduitHandler, c *routes.HandlerContext) httprouter.Handle {
	middleware := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		next(w, r, ps, c)
	}

	return middleware
}

func ConduitAuthMiddleware(next routes.AuthHandler, c *routes.HandlerContext) httprouter.Handle {

	middleware := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		routes.SetCorsHeaders(w)
		if candidateToken, ok := r.Header["X-Access-Token"]; ok {
			// Parse and validate token:
			token, err := jwt.ParseWithClaims(candidateToken[0], &routes.HomeAutoClaims{}, func(token *jwt.Token) (interface{}, error) {
				return routes.SecretKey, nil
			})

			if claims, ok := token.Claims.(*routes.HomeAutoClaims); ok && token.Valid {
				next(w, r, ps, c, claims)
				return
			} else {
				routes.SendErrorResponse(w, err.Error(), 401)
				fmt.Println("Error in Auth middleware")
				fmt.Println(err.Error())
				return
			}
		}
		// Either token wasn't valid or it wasn't provided
		routes.SendErrorResponse(w, "No Token", 400)
		return
	}

	return middleware

}
